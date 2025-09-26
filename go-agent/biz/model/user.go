package model


import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(50);uniqueIndex;"`
	Email    string `json:"email" gorm:"type:varchar(50);uniqueIndex;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	Nickname string `json:"nickname" gorm:"type:varchar(50)"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255)"`
}