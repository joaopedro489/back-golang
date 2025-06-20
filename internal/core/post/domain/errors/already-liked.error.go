package errors

type AlreadyLikedError struct {
	Message string
}

func NewAlreadyLikedError() *AlreadyLikedError {
	return &AlreadyLikedError{
		Message: "You have already liked this post.",
	}
}

func (e *AlreadyLikedError) Error() string {
	return e.Message
}
