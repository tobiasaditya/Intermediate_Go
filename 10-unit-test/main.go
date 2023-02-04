package main

import (
	"10-unit-test/repository"
	"errors"
	"log"
)

func main() {

}

func LuasPersegi(sisi int) int {
	return sisi * sisi
}

func Register(username string, password string) error {
	if username == "" {
		return errors.New("username empty")
	}

	if password == "" {
		return errors.New("password empty")
	}

	return nil
}

func RegisterToDB(userRepo repository.IUserRepository, username string, password string) error {
	if err := userRepo.Register(username, password); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
