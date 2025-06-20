package errors

type UserNotFoundError struct {
	Message string
}

func NewUserNotFoundError() *UserNotFoundError {
	return &UserNotFoundError{
		Message: "Post not found",
	}
}

func (e *UserNotFoundError) Error() string {
	return e.Message
}
