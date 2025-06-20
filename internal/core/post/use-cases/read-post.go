package usecases

import (
	"github.com/joaopedro489/back-golang/internal/core/post/domain/errors"
	"github.com/joaopedro489/back-golang/internal/core/post/domain/repository"
	"github.com/joaopedro489/back-golang/internal/core/post/types/input"
	"github.com/joaopedro489/back-golang/internal/core/post/types/output"
)

type ReadPost struct {
	repo     repository.IPostRepository
	userRepo repository.IUserRepository
}

func NewReadPost(repo repository.IPostRepository, userRepo repository.IUserRepository) *ReadPost {
	return &ReadPost{repo: repo, userRepo: userRepo}
}

func (uc *ReadPost) Execute(params input.ParamPost) (*output.PostResponse, error) {
	post, err := uc.repo.GetPostById(params.Id)
	if err != nil {
		return nil, err
	}

	if post == nil {
		postErrors := errors.NewPostErrors()
		return nil, postErrors.PostNotFoundError
	}

	user, err := uc.userRepo.GetUserName(post.UserId)
	if err != nil {
		return nil, err
	}

	if user == "" {
		postErrors := errors.NewPostErrors()
		return nil, postErrors.UserNotFoundError
	}

	postResponse := output.NewPostResponse(post.Id, post.Title, post.Content, user, len(post.Likes))

	return postResponse, nil
}
