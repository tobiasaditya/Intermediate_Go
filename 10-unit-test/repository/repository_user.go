package repository

import "fmt"

type IUserRepository interface {
	Register(username string, password string) error
}

type UserRepository struct{}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Register(username string, password string) error {
	fmt.Println("Inserting to db")
	return nil
}
