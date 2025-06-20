package usecases

import (
	"github.com/joaopedro489/back-golang/internal/core/users/domain"
	"github.com/joaopedro489/back-golang/internal/core/users/domain/errors"
	"github.com/joaopedro489/back-golang/internal/core/users/types/input"
)

type ReadUser struct {
	repo domain.IUserRepository
}

func NewReadUser(repo domain.IUserRepository) *ReadUser {
	return &ReadUser{repo: repo}
}

func (uc *ReadUser) Execute(params input.ParamUser) (*domain.User, error) {
	user, err := uc.repo.GetUserById(params.Id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		usersErrors := errors.NewUserErrors()
		return nil, usersErrors.UserNotFoundError
	}

	return user, nil
}
