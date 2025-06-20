package usecases

import (
	"github.com/joaopedro489/back-golang/internal/core/post/domain/errors"
	"github.com/joaopedro489/back-golang/internal/core/post/domain/repository"
)

type LikePost struct {
	repo repository.IPostRepository
}

func NewLikePost(repo repository.IPostRepository) *LikePost {
	return &LikePost{repo: repo}
}

func (uc *LikePost) Execute(postId int, userId int) error {
	post, err := uc.repo.GetPostById(postId)
	if err != nil {
		return err
	}

	if post == nil {
		postErrors := errors.NewPostErrors()
		return postErrors.PostNotFoundError
	}

	alreadyLiked, err := uc.repo.CheckAlreadyLiked(postId, userId)
	if err != nil {
		return err
	}
	if alreadyLiked {
		postErrors := errors.NewPostErrors()
		return postErrors.AlreadyLikedError
	}

	if err := uc.repo.LikePost(postId, userId); err != nil {
		return err
	}

	return nil
}
