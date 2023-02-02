package service

import (
	"9-session-login/dto"
	"9-session-login/entity"
	"9-session-login/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) RegisterUser(inputUser dto.InputUser) (entity.User, error) {
	newUser, err := s.userRepository.CreateUser(inputUser)
	if err != nil {
		return entity.User{}, err
	}
	return newUser, nil
}

func (s *UserService) LoginUser(inputLogin dto.InputLogin) (entity.User, error) {
	loggedUser, err := s.userRepository.Login(inputLogin.UserName, inputLogin.Password)
	if err != nil {
		return entity.User{}, err
	}
	return loggedUser, err
}
