package repository

import (
	"9-session-login/dto"
	"9-session-login/entity"

	"database/sql"
)

// type userRepository interface {
// 	CreateUser(newUser entity.User) (entity.User, error)
// 	Login(username string, password string) (entity.User, error)
// }

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(newUser dto.InputUser) (entity.User, error) {
	var userId int
	err := r.db.QueryRow(`
	INSERT INTO users(username,first_name,last_name,password)
	VALUES($1,$2,$3,$4) RETURNING id`, newUser.UserName, newUser.FirstName, newUser.LastName, newUser.Password).Scan(&userId)
	if err != nil {
		return entity.User{}, err
	}

	insertedUser := entity.User{
		ID:        userId,
		UserName:  newUser.UserName,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Password:  newUser.Password,
	}

	return insertedUser, nil
}

func (r *UserRepository) Login(username string, password string) (entity.User, error) {
	var foundUser entity.User
	err := r.db.QueryRow(
		`SELECT * FROM users WHERE username = $1 AND password = $2`, username, password).Scan(&foundUser.ID, &foundUser.UserName, &foundUser.FirstName, &foundUser.LastName, &foundUser.Password)
	if err != nil {
		return entity.User{}, err
	}

	return foundUser, nil
}
