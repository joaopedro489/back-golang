package usecases

import (
	"github.com/joaopedro489/back-golang/internal/core/post/domain/errors"
	"github.com/joaopedro489/back-golang/internal/core/post/domain/policies"
	"github.com/joaopedro489/back-golang/internal/core/post/domain/repository"
	"github.com/joaopedro489/back-golang/internal/core/post/types/input"
)

type DeletePost struct {
	repo repository.IPostRepository
}

func NewDeletePost(repo repository.IPostRepository) *DeletePost {
	return &DeletePost{repo: repo}
}

func (uc *DeletePost) Execute(params input.ParamPost, userId int) error {
	post, err := uc.repo.GetPostById(params.Id)
	if err != nil {
		return err
	}

	if post == nil {
		postErrors := errors.NewPostErrors()
		return postErrors.PostNotFoundError
	}

	if err := policies.CanDeletePost(post, userId); err != nil {
		return err
	}

	if err := uc.repo.DeletePost(post.Id); err != nil {
		return err
	}

	return nil
}
