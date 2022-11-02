package createUser

import (
	model "github.com/053steve/gin-boilerplate/models"
)

type Service interface {
	CreateStudentService(input *InputCreateStudent) (*model.EntityStudent, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateStudentService(input *InputCreateStudent) (*model.EntityStudent, string) {

	students := model.EntityStudent{
		Name: input.Name,
		Npm:  input.Npm,
		Fak:  input.Fak,
		Bid:  input.Bid,
	}

	resultCreateStudent, errCreateStudent := s.repository.CreateStudentRepository(&students)

	return resultCreateStudent, errCreateStudent
}

func (s *service) CreateUserService(input *InputCreateStudent) (*model.EntityUsers, string) {

}

// type EntityUsers struct {
// 	ID        string `gorm:"primaryKey;"`
// 	UserName  string `gorm:"type:varchar(255);not null"`
// 	Password  string `gorm:"type:varchar(255);not null"`
// 	FirstName string `gorm:"type:varchar(255);not null"`
// 	LastName  string `gorm:"type:varchar(255);not null"`
// 	UserType  string `gorm:"type:varchar(100);not null"`
// 	Email     string `gorm:"type:varchar(255);unique;not null"`
// 	Active    bool   `gorm:"type:bool;default:false"`
// 	PublicKey bool   `gorm:"type:varchar(255);not null"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }
