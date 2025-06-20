package usecases

import (
	"github.com/joaopedro489/back-golang/internal/core/users/domain"
	"github.com/joaopedro489/back-golang/internal/core/users/domain/errors"
	"github.com/joaopedro489/back-golang/internal/core/users/types/input"
)

type EditUser struct {
	repo domain.IUserRepository
}

func NewEditUser(repo domain.IUserRepository) *EditUser {
	return &EditUser{repo: repo}
}

func (uc *EditUser) Execute(params input.EditUser, id input.ParamUser) (*domain.User, error) {
	if err := input.ValidateEditUser(params); err != nil {
		return nil, err
	}

	user, err := uc.repo.GetUserById(id.Id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		usersErrors := errors.NewUserErrors()
		return nil, usersErrors.UserNotFoundError
	}

	user.Replace(params)

	err = uc.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
