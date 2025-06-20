package errors

type UserNotFoundError struct {
	Message string
}

func NewUserNotFoundError() *UserNotFoundError {
	return &UserNotFoundError{
		Message: "User not found",
	}
}

func (e *UserNotFoundError) Error() string {
	return e.Message
}
