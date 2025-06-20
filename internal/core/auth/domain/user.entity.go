package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id    int
	Email string
	Hash  string
}

func NewUser(user User) *User {
	return &User{
		Id:    user.Id,
		Email: user.Email,
		Hash:  user.Hash,
	}
}

func (user *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
	return err == nil
}
