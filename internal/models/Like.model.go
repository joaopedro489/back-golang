package models

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	PostId int `gorm:"not null"`
	UserId int `gorm:"not null"`
}
