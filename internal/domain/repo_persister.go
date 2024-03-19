package domain

type RepoPersister interface {
	CreateUser(User) (User, error)
	UpdateUser(User) (User, error)
	ListUsers(User) ([]User, error)
	DeleteUser(string) error
}
