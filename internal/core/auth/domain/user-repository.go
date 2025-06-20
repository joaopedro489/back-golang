package domain

type IUserRepository interface {
	GetByEmail(email string) (*User, error)
}
