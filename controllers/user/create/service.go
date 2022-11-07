package createUser

import (
	model "github.com/053steve/gin-boilerplate/models"
)

type Service interface {
	CreateUserService(input *CreateUserDto) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateUserService(input *CreateUserDto) (*model.User, string) {
	user := model.User{
		UserName:  input.UserName,
		Email:     input.Email,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		UserType:  input.UserType,
		Active:    input.Active,
	}

	resultCreateUser, errCreateUser := s.repository.CreateStudentRepository(&user)

	return resultCreateUser, errCreateUser
}
