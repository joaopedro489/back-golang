package errors

type UserErrors struct {
	UserNotFoundError *UserNotFoundError
}

func NewUserErrors() *UserErrors {
	return &UserErrors{
		UserNotFoundError: NewUserNotFoundError(),
	}
}
