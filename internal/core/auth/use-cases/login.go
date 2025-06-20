package usecases

import (
	"github.com/joaopedro489/back-golang/internal/auth"
	"github.com/joaopedro489/back-golang/internal/core/auth/domain"
	"github.com/joaopedro489/back-golang/internal/core/auth/domain/errors"
	"github.com/joaopedro489/back-golang/internal/core/auth/types/input"
	"github.com/joaopedro489/back-golang/internal/core/auth/types/output"
)

type Login struct {
	repo domain.IUserRepository
}

func NewLogin(repo domain.IUserRepository) *Login {
	return &Login{repo: repo}
}

func (uc *Login) Execute(params input.LoginInput) (*output.TokenResponse, error) {
	if err := input.ValidateLoginInput(params); err != nil {
		return nil, err
	}

	user, err := uc.repo.GetByEmail(params.Email)
	if err != nil {
		return nil, err
	}

	if user == nil || !user.CheckPasswordHash(params.Password) {
		authErrors := errors.NewAuthErrors()
		return nil, authErrors.InvalidPasswordError
	}

	token, err := auth.GenerateJWT(user.Id)

	if err != nil {
		return nil, err
	}

	return &output.TokenResponse{
		Token: token,
	}, nil
}
