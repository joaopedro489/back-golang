package usecases

import (
	"github.com/joaopedro489/back-golang/internal/core/post/domain/entity"
	"github.com/joaopedro489/back-golang/internal/core/post/domain/repository"
	"github.com/joaopedro489/back-golang/internal/core/post/types/input"
)

type BrowsePost struct {
	repo repository.IPostRepository
}

func NewBrowsePost(repo repository.IPostRepository) *BrowsePost {
	return &BrowsePost{repo: repo}
}

func (uc *BrowsePost) Execute(params input.QueryPost) ([]entity.Post, error) {
	filters := repository.Filters{
		Limit:  10,
		Offset: 0,
		UserId: params.UserId,
	}
	if params.Limit != nil {
		filters.Limit = *params.Limit
	}
	if params.Offset != nil {
		filters.Offset = *params.Offset
	}
	posts, err := uc.repo.BrowsePosts(filters)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
