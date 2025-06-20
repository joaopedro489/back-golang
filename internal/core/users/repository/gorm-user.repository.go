package repository

import (
	"github.com/joaopedro489/back-golang/internal/core/users/domain"
	"github.com/joaopedro489/back-golang/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) BrowseUsers(filters domain.Filters) ([]domain.User, error) {
	var users []models.User
	query := r.db.Model(&models.User{})

	if filters.Search != nil {
		query = query.Where("name ILIKE ?", "%"+*filters.Search+"%")
	}

	query = query.Limit(filters.Limit)
	query = query.Offset(filters.Offset)

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	userDomains := domain.UserModelsToDomains(users)
	return userDomains, nil
}

func (r *UserRepository) CreateUser(user domain.User) (*domain.User, error) {
	dbUser := domain.UserDomainToModel(user)
	if err := r.db.Create(&dbUser).Error; err != nil {
		return nil, err
	}
	createdUser := domain.UserModelToDomain(dbUser)
	return &createdUser, nil
}

func (r *UserRepository) GetUserById(id int) (*domain.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	userDomain := domain.UserModelToDomain(user)
	return &userDomain, nil
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
	dbUser := domain.UserDomainToModel(*user)
	if err := r.db.Model(&dbUser).Updates(dbUser).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
