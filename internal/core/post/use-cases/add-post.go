package usecases

import (
	"github.com/joaopedro489/back-golang/internal/core/post/domain/entity"
	"github.com/joaopedro489/back-golang/internal/core/post/domain/errors"
	"github.com/joaopedro489/back-golang/internal/core/post/domain/repository"
	"github.com/joaopedro489/back-golang/internal/core/post/types/input"
	"github.com/joaopedro489/back-golang/internal/core/post/types/output"
)

type AddPost struct {
	repo     repository.IPostRepository
	userRepo repository.IUserRepository
}

func NewAddPost(repo repository.IPostRepository, userRepo repository.IUserRepository) *AddPost {
	return &AddPost{repo: repo, userRepo: userRepo}
}

func (uc *AddPost) Execute(params input.AddPost, userId int) (*output.PostResponse, error) {
	if err := input.ValidateAddPost(params); err != nil {
		return nil, err
	}

	postChan := make(chan repository.PostResult, 1)
	userChan := make(chan repository.UserResult, 1)

	go uc.sendPostData(params, userId, postChan)
	go uc.sendUserData(userId, userChan)

	createdPost, userName, err := uc.getResults(postChan, userChan)
	if err != nil {
		return nil, err
	}

	if userName == "" {
		postErrors := errors.NewPostErrors()
		return nil, postErrors.UserNotFoundError
	}

	postResponse := output.NewPostResponse(createdPost.Id, createdPost.Title, createdPost.Content, userName, len(createdPost.Likes))

	return postResponse, nil
}

func (uc *AddPost) sendPostData(params input.AddPost, userId int, postChan chan<- repository.PostResult) {
	post := entity.NewPost(params.Title, params.Content, userId)
	createdPost, err := uc.repo.CreatePost(*post)
	postChan <- repository.PostResult{Post: createdPost, Err: err}
	close(postChan)
}

func (uc *AddPost) sendUserData(userId int, userChan chan<- repository.UserResult) {
	name, err := uc.userRepo.GetUserName(userId)
	userChan <- repository.UserResult{Name: name, Err: err}
	close(userChan)
}

func (uc *AddPost) getResults(postChan <-chan repository.PostResult, userChan <-chan repository.UserResult) (*entity.Post, string, error) {
	var createdPost *entity.Post
	var userName string
	var postErr, userErr error
	postReceived := false
	userReceived := false

	for !postReceived || !userReceived {
		select {
		case pr, ok := <-postChan:
			if ok {
				createdPost = pr.Post
				postErr = pr.Err
				postReceived = true
			}
		case ur, ok := <-userChan:
			if ok {
				userName = ur.Name
				userErr = ur.Err
				userReceived = true
			}
		}
	}

	if postErr != nil {
		return nil, "", postErr
	}
	if userErr != nil {
		return nil, "", userErr
	}

	return createdPost, userName, nil
}
