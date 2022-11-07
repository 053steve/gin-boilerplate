package model

import (
	"time"

	util "github.com/053steve/gin-boilerplate/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntityUser struct {
	ID        string `gorm:"primaryKey;"`
	UserName  string `gorm:"type:varchar(255);not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName  string `gorm:"type:varchar(255);not null"`
	UserType  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Active    bool   `gorm:"type:bool;default:false"`
	PublicKey bool   `gorm:"type:varchar(255);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *EntityUser) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.Password = util.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityUser) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
