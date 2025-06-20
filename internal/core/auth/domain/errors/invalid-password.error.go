package errors

type InvalidPasswordError struct {
	Message string
}

func NewInvalidPasswordError() *InvalidPasswordError {
	return &InvalidPasswordError{
		Message: "Invalid password",
	}
}

func (e *InvalidPasswordError) Error() string {
	return e.Message
}
