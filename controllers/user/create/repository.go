package createUser

import (
	model "github.com/053steve/gin-boilerplate/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateStudentRepository(input *model.EntityUser) (*model.EntityUser, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateStudentRepository(input *model.EntityUser) (*model.EntityUser, string) {

	var user model.EntityUser
	db := r.db.Model(&user)
	errorCode := make(chan string, 1)

	checkStudentExist := db.Debug().Select("*").Where("email = ?", input.Email).Find(&user)

	if checkStudentExist.RowsAffected > 0 {
		errorCode <- "CREATE_STUDENT_CONFLICT_409"
		return &user, <-errorCode
	}

	user = *input

	addNewUser := db.Debug().Create(&user)
	db.Commit()

	if addNewUser.Error != nil {
		errorCode <- "CREATE_STUDENT_FAILED_403"
		return &user, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &user, <-errorCode
}
