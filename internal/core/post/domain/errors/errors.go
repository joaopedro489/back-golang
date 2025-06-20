package errors

type PostErrors struct {
	PostNotFoundError     *PostNotFoundError
	UserNotFoundError     *UserNotFoundError
	CannotDeletePostError *CannotDeletePostError
	AlreadyLikedError     *AlreadyLikedError
}

func NewPostErrors() *PostErrors {
	return &PostErrors{
		PostNotFoundError:     NewPostNotFoundError(),
		UserNotFoundError:     NewUserNotFoundError(),
		CannotDeletePostError: NewCannotDeletePostError(),
		AlreadyLikedError:     NewAlreadyLikedError(),
	}
}
