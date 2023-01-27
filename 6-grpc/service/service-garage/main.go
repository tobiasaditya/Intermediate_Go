package main

import (
	"6-grpc/common/config"
	"6-grpc/common/model"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var localStorage *model.GarageListByUser

func init() {
	localStorage = new(model.GarageListByUser)
	localStorage.List = make(map[string]*model.GarageList)
}

func main() {
	srv := grpc.NewServer()
	var garageSrv GaragesServer
	model.RegisterGaragesServer(srv, garageSrv)

	l, err := net.Listen("tcp", config.SERVICE_GARAGE_PORT)
	if err != nil {
		log.Fatal("error listening to port:", config.SERVICE_GARAGE_PORT)
	}

	// setup http proxy
	go func() {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		grpcServerEndpoint := flag.String("grpc-server-endpoint", "localhost"+config.SERVICE_USER_PORT, "gRPC server endpoint")
		_ = model.RegisterUsersHandlerFromEndpoint(context.Background(), mux, *grpcServerEndpoint, opts)
		log.Println("Starting User Server HTTP at 9001 ")
		http.ListenAndServe(":9001", mux)
	}()

	srv.Serve(l)

}

type GaragesServer struct {
	model.UnimplementedGaragesServer
}

func (gs GaragesServer) List(ctx context.Context, param *model.GarageUserId) (*model.GarageList, error) {
	userId := param.UserId
	data, ok := localStorage.List[userId]
	if !ok {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	return data, nil
}
func (gs GaragesServer) Add(ctx context.Context, param *model.GarageAndUserId) (*emptypb.Empty, error) {
	fmt.Println("inserting garage")
	userId := param.UserId
	garage := param.Garage

	//Try getting element by userId
	_, ok := localStorage.List[userId]

	//If userId not found on list, make a new one
	if !ok {
		localStorage.List[userId] = new(model.GarageList)
		localStorage.List[userId].List = make([]*model.Garage, 0)
	}

	//Append garage to corresponding userId
	localStorage.List[userId].List = append(localStorage.List[userId].List, garage)

	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
