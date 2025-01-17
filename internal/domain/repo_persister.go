package domain

type Repository interface {
	CreateUser(User) (User, error)
	UpdateUser(User) (User, error)
	ListUsers([]Param) ([]User, error)
	GetUser(string) (User, error)
	DeleteUser(string) error

	CreateService(Service, string, string) (Service, error)
	ListServices(params []Param) ([]Service, error)
}
