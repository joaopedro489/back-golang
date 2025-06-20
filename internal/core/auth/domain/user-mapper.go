package domain

import (
	"github.com/joaopedro489/back-golang/internal/models"
)

func UserModelToDomain(m models.User) User {
	return User{
		Id:    int(m.ID),
		Email: m.Email,
		Hash:  m.Hash,
	}
}
