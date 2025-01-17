package domain

type Actions interface {
	CreateUser(User) (User, error)
	UpdateUser(User) (User, error)
	ListUsers([]Param) ([]User, error)
	DeleteUser(string) error

	CreateService(service Service) (Service, error)
	ListServices(params []Param) ([]Service, error)
}

type actions struct {
	db Repository
}

func ProviderService(db Repository) Actions {
	return &actions{
		db: db,
	}
}
