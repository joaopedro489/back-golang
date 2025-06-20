package input

import (
	"github.com/go-playground/validator/v10"
)

type AddPost struct {
	Title   string `validate:"required,min=3"`
	Content string `validate:"required,min=3"`
}

func ValidateAddPost(post AddPost) error {
	validate := validator.New()
	return validate.Struct(post)
}
