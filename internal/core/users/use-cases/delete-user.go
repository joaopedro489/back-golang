package usecases

import (
	"github.com/joaopedro489/back-golang/internal/core/users/domain"
	"github.com/joaopedro489/back-golang/internal/core/users/domain/errors"
	"github.com/joaopedro489/back-golang/internal/core/users/types/input"
)

type DeleteUser struct {
	repo domain.IUserRepository
}

func NewDeleteUser(repo domain.IUserRepository) *DeleteUser {
	return &DeleteUser{repo: repo}
}

func (uc *DeleteUser) Execute(params input.ParamUser) error {
	user, err := uc.repo.GetUserById(params.Id)
	if err != nil {
		return err
	}

	if user == nil {
		usersErrors := errors.NewUserErrors()
		return usersErrors.UserNotFoundError
	}

	if err := uc.repo.DeleteUser(user.Id); err != nil {
		return err
	}

	return nil
}
