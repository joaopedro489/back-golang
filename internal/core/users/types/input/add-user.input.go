package input

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type AddUser struct {
	Name     string    `validate:"required,min=3"`
	Email    string    `validate:"required,email"`
	Birthday time.Time `validate:"required,birthday_past"`
	Password string    `validate:"required,min=6"`
}

func ValidateAddUser(user AddUser) error {
	validate := validator.New()
	validate.RegisterValidation("birthday_past", func(fl validator.FieldLevel) bool {
		birthday := fl.Field().Interface().(time.Time)
		return birthday.Before(time.Now())
	})
	return validate.Struct(user)
}
