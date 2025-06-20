package policies

import (
	"github.com/joaopedro489/back-golang/internal/core/post/domain/entity"
	"github.com/joaopedro489/back-golang/internal/core/post/domain/errors"
)

func CanDeletePost(post *entity.Post, userId int) error {
	if post.UserId != userId {
		postErrors := errors.NewPostErrors()
		return postErrors.CannotDeletePostError
	}
	return nil
}
