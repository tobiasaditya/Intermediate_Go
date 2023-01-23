package main

import (
	"6-grpc/common/config"
	"6-grpc/common/model"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to ", port, err)
	}
	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to ", port, err)
	}
	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "1",
		Name:     "Tobias",
		Password: "password",
		Gender:   model.UserGender_MALE,
	}

	user2 := model.User{
		Id:       "2",
		Name:     "Aditya",
		Password: "password",
		Gender:   model.UserGender_UNDEFINED,
	}

	// garage1 := model.Garage{
	// 	Id:   "1",
	// 	Name: "Garage One",
	// 	Coordinate: &model.GarageCoordinate{
	// 		Latitude:  1.0,
	// 		Longitude: 32.0,
	// 	},
	// }

	userService := serviceUser()
	fmt.Println("=====USER SERVICE TEST======")
	userService.Register(context.Background(), &user1)
	userService.Register(context.Background(), &user2)
	data, err := userService.List(context.Background(), new(emptypb.Empty))

	if err != nil {
		log.Fatalln(err.Error())
	}

	resString, err := json.Marshal(data.List)

	log.Println(string(resString))

}
