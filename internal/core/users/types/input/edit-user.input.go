package input

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type EditUser struct {
	Name     *string    `validate:"omitempty,min=3"`
	Email    *string    `validate:"omitempty,email"`
	Birthday *time.Time `validate:"omitempty,birthday_past"`
}

func ValidateEditUser(user EditUser) error {
	validate := validator.New()
	validate.RegisterValidation("birthday_past", func(fl validator.FieldLevel) bool {
		if birthday, ok := fl.Field().Interface().(*time.Time); ok && birthday != nil {
			return (*birthday).Before(time.Now())
		}
		return true
	})
	return validate.Struct(user)
}
