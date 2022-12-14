package registerAuth

import (
	model "github.com/053steve/gin-boilerplate/models"
)

type Service interface {
	RegisterService(input *RegisterDto) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceRegister(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterService(input *RegisterDto) (*model.User, string) {

	users := model.User{
		Email:     input.Email,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		UserType:  input.UserType,
	}

	resultRegister, errRegister := s.repository.RegisterRepository(&users)

	return resultRegister, errRegister
}
