package domain

type Filters struct {
	Search *string
	Limit  int
	Offset int
}

type IUserRepository interface {
	GetUserById(id int) (*User, error)
	BrowseUsers(filters Filters) ([]User, error)
	CreateUser(user User) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error
}
