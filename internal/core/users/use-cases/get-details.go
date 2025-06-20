package usecases

import (
	"github.com/joaopedro489/back-golang/internal/core/users/domain"
	"github.com/joaopedro489/back-golang/internal/core/users/domain/errors"
)

type GetDetails struct {
	repo domain.IUserRepository
}

func NewGetDetails(repo domain.IUserRepository) *GetDetails {
	return &GetDetails{repo: repo}
}

func (uc *GetDetails) Execute(userId int) (*domain.User, error) {
	user, err := uc.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		usersErrors := errors.NewUserErrors()
		return nil, usersErrors.UserNotFoundError
	}

	return user, nil
}
