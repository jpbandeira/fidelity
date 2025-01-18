package domain

type Actions interface {
	CreateClient(Client) (Client, error)
	UpdateClient(Client) (Client, error)
	ListClients([]Param) ([]Client, error)
	DeleteClient(string) error

	CreateAttendant(Attendant) (Attendant, error)
	UpdateAttendant(Attendant) (Attendant, error)
	ListAttendants([]Param) ([]Attendant, error)
	DeleteAttendant(string) error

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
