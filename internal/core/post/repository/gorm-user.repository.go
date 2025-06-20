package repository

import (
	"github.com/joaopedro489/back-golang/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserName(id int) (string, error) {
	var user models.User
	if err := r.db.Model(&user).Select("name").Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", err
	}
	return user.Name, nil
}
