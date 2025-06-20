package usecases

import (
	"github.com/joaopedro489/back-golang/internal/core/users/domain"
	types "github.com/joaopedro489/back-golang/internal/core/users/types/input"
	"golang.org/x/crypto/bcrypt"
)

type AddUser struct {
	repo domain.IUserRepository
}

func NewAddUser(repo domain.IUserRepository) *AddUser {
	return &AddUser{repo: repo}
}

func (uc *AddUser) Execute(params types.AddUser) (*domain.User, error) {
	if err := types.ValidateAddUser(params); err != nil {
		return nil, err
	}

	hash, err := HashPassword(params.Password)
	if err != nil {
		return nil, err
	}

	user := domain.User{
		Name:     params.Name,
		Email:    params.Email,
		Birthday: params.Birthday,
		Hash:     hash,
	}

	createdUser, err := uc.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
