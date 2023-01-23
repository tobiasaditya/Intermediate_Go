package main

import (
	"6-grpc/common/config"
	"6-grpc/common/model"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)

}

func main() {
	srv := grpc.NewServer()
	var userSrv UsersServer
	model.RegisterUsersServer(srv, userSrv)
	l, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatal("error listening to port:", config.SERVICE_USER_PORT)
	}

	srv.Serve(l)
}

type UsersServer struct {
	model.UnimplementedUsersServer
}

func (us UsersServer) Register(ctx context.Context, user *model.User) (*emptypb.Empty, error) {
	localStorage.List = append(localStorage.List, user)
	log.Println("inserted user")
	return new(emptypb.Empty), nil
}
func (us UsersServer) List(context.Context, *emptypb.Empty) (*model.UserList, error) {
	return localStorage, nil
}
