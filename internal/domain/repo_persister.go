package domain

type RepoPersister interface {
	CreateUser(User) (User, error)
	UpdateUser(User) (User, error)
	ListUsers([]Param) ([]User, error)
	DeleteUser(string) error

	CreateService(service Service) (Service, error)
}
