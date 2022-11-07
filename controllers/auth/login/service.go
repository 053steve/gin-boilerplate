package loginAuth

import (
	model "github.com/053steve/gin-boilerplate/models"
)

type Service interface {
	LoginService(input *LoginDto) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginService(input *LoginDto) (*model.User, string) {

	user := model.User{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.LoginRepository(&user)

	return resultLogin, errLogin
}
