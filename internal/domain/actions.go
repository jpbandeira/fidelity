package domain

type Actions interface {
	CreateClient(Client) (Client, error)
	UpdateClient(Client) (Client, error)
	ListClients([]Param) ([]Client, error)
	DeleteClient(string) error

	CreateAppointment(appt Appointment) (Appointment, error)
	ListServices([]Param) ([]Service, error)
	GetServiceSummary(string) ([]ServiceSummary, error)

	ListServiceTypes([]Param) ([]ServiceType, error)
	CreateServiceType(st ServiceType) (ServiceType, error)
}

type actions struct {
	db Repository
}

func ProviderService(db Repository) Actions {
	return &actions{
		db: db,
	}
}
