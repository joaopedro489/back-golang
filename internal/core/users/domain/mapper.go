package domain

import (
	"github.com/joaopedro489/back-golang/internal/models"
	"gorm.io/gorm"
)

func UserModelToDomain(m models.User) User {
	return User{
		Id:       int(m.ID),
		Name:     m.Name,
		Email:    m.Email,
		Birthday: m.Birthday,
	}
}

func UserModelsToDomains(models []models.User) []User {
	userDomains := make([]User, len(models))
	for i, user := range models {
		userDomains[i] = UserModelToDomain(user)
	}
	return userDomains
}

func UserDomainToModel(domain User) models.User {
	return models.User{
		Model: gorm.Model{
			ID: uint(domain.Id),
		},
		Name:     domain.Name,
		Email:    domain.Email,
		Birthday: domain.Birthday,
		Hash:     domain.Hash,
	}
}
