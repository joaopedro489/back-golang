package input

import "github.com/go-playground/validator/v10"

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func ValidateLoginInput(input LoginInput) error {
	validate := validator.New()
	return validate.Struct(input)
}
