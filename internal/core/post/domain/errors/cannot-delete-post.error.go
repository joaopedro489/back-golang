package errors

type CannotDeletePostError struct {
	Message string
}

func NewCannotDeletePostError() *CannotDeletePostError {
	return &CannotDeletePostError{
		Message: "Cannot delete Post",
	}
}

func (e *CannotDeletePostError) Error() string {
	return e.Message
}
