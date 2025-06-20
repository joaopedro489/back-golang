package repository

type UserResult struct {
	Name string
	Err  error
}

type IUserRepository interface {
	GetUserName(id int) (string, error)
}
