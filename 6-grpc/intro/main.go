package main

import (
	"6-grpc/model"
	"bytes"
	"fmt"
	"os"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	user1 := &model.User{
		Id:       "u001",
		Name:     "Tobias Aditya",
		Password: "password",
		Gender:   model.UserGender_MALE,
	}

	// userList := &model.UserList{
	// 	List: []*model.User{
	// 		user1,
	// 	},
	// }

	garage1 := &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.5,
			Longitude: 53.2,
		},
	}

	garageList := &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	// garageListByUser := &model.GarageListByUser{
	// 	List: map[string]*model.GarageList{
	// 		user1.Id: garageList,
	// 	},
	// }

	fmt.Printf("# ==== Original\n    %#v \n", user1)
	fmt.Printf("# ====As String\n   %v\n", user1.String())

	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)

	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}

	jsonString := buf.String()
	fmt.Println("JSON String : ", jsonString)

	protoObject := new(model.GarageList)

	//Using jsonpb.Unmarshaler
	// buf2 := strings.NewReader(jsonString)
	// err2 := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject)
	// if err2 != nil {
	// 	fmt.Println(err2.Error())
	// 	os.Exit(0)
	// }

	// Using jsonpb.UnmarshalString
	err2 := jsonpb.UnmarshalString(jsonString, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}
	fmt.Println("As String : ", protoObject.String())

}
