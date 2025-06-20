package domain

import (
	"time"

	"github.com/joaopedro489/back-golang/internal/core/users/types/input"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Birthday time.Time
	Hash     string
}

func NewUser(user User) *User {
	return &User{
		Name:     user.Name,
		Email:    user.Email,
		Birthday: user.Birthday,
		Id:       user.Id,
		Hash:     user.Hash,
	}
}

func (user *User) Replace(edit input.EditUser) {
	if edit.Name != nil {
		user.Name = *edit.Name
	}
	if edit.Email != nil {
		user.Email = *edit.Email
	}
	if edit.Birthday != nil {
		user.Birthday = *edit.Birthday
	}
}
