package domain

type RepoPersister interface {
	CreateUser(User) (User, error)
	UpdateUser(User) (User, error)
	ListUsers([]Param) ([]User, error)
	GetUser(string) (User, error)
	DeleteUser(string) error

	CreateService(Service, string, string) (Service, error)
}
