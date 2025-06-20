package errors

type PostNotFoundError struct {
	Message string
}

func NewPostNotFoundError() *PostNotFoundError {
	return &PostNotFoundError{
		Message: "Post not found",
	}
}

func (e *PostNotFoundError) Error() string {
	return e.Message
}
