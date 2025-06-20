package usecases

import (
	"github.com/joaopedro489/back-golang/internal/core/users/domain"
	"github.com/joaopedro489/back-golang/internal/core/users/types/input"
)

type BrowseUser struct {
	repo domain.IUserRepository
}

func NewBrowseUser(repo domain.IUserRepository) *BrowseUser {
	return &BrowseUser{repo: repo}
}

func (uc *BrowseUser) Execute(params input.QueryUser) ([]domain.User, error) {
	filters := domain.Filters{
		Limit:  10,
		Offset: 0,
		Search: params.Search,
	}
	if params.Limit != nil {
		filters.Limit = *params.Limit
	}
	if params.Offset != nil {
		filters.Offset = *params.Offset
	}

	users, err := uc.repo.BrowseUsers(filters)
	if err != nil {
		return nil, err
	}

	return users, nil
}
