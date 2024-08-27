package domain

type Actions interface {
	CreateUser(User) (User, error)
	UpdateUser(User) (User, error)
	ListUsers([]Param) ([]User, error)
	DeleteUser(string) error

	CreateService(service Service) (Service, error)
}

type actions struct {
	db RepoPersister
}

func ProviderService(db RepoPersister) Actions {
	return &actions{
		db: db,
	}
}
