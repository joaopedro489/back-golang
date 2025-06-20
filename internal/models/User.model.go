package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"not null;unique"`
	Birthday time.Time
	Hash     string `gorm:"not null"`
	Posts    []Post
	Likes    []Like
}
