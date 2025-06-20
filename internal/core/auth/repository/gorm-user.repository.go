package repository

import (
	"github.com/joaopedro489/back-golang/internal/core/auth/domain"
	"github.com/joaopedro489/back-golang/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	foundUser := domain.UserModelToDomain(user)
	return &foundUser, nil
}
