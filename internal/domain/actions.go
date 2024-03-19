package domain

type Actions interface {
	CreateUser(User) (User, error)
	UpdateUser(User) (User, error)
	ListUsers() ([]User, error)
	DeleteUser(string) error
}

type actions struct {
	db RepoPersister
}

func ProviderService(db RepoPersister) Actions {
	return &actions{
		db: db,
	}
}
