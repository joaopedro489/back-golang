package errors

type AuthErrors struct {
	InvalidPasswordError *InvalidPasswordError
}

func NewAuthErrors() *AuthErrors {
	return &AuthErrors{
		InvalidPasswordError: NewInvalidPasswordError(),
	}
}
