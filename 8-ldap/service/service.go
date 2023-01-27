package service

import (
	"8-ldap/model"
	"8-ldap/repository"
	"errors"

	"github.com/labstack/gommon/log"
)

type LoginService struct {
	repository *repository.LDAPRepo
}

func NewLoginService(repository *repository.LDAPRepo) *LoginService {
	return &LoginService{repository: repository}
}

func (s *LoginService) Login(username string, password string) (*model.UserLDAPData, error) {
	ok, data, err := s.repository.AuthUsingLDAP(username, password)
	if !ok {
		err := errors.New("auth using ldap not ok")
		log.Error("auth using ldap not ok")
		return nil, err
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return data, nil
}
