package output

import (
	"time"

	"github.com/joaopedro489/back-golang/internal/core/users/domain"
)

type UserResponse struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
}

func NewUserResponse(name, email string, birthday time.Time, id int) *UserResponse {
	return &UserResponse{
		Id:       id,
		Name:     name,
		Email:    email,
		Birthday: birthday,
	}
}

func NewUsersResponse(users []domain.User) []UserResponse {
	var responses []UserResponse
	for _, user := range users {
		userResponse := NewUserResponse(user.Name, user.Email, user.Birthday, user.Id)
		responses = append(responses, *userResponse)
	}
	return responses
}
