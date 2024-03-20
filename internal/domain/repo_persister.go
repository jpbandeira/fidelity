package domain

type RepoPersister interface {
	CreateUser(User) (User, error)
	UpdateUser(User) (User, error)
	ListUsers() ([]User, error)
	DeleteUser(string) error

	CreateService(service Service) (Service, error)
}
