package domain

type Repository interface {
	CreateClient(Client) (Client, error)
	UpdateClient(Client) (Client, error)
	ListClients([]Param) ([]Client, error)
	GetClient(string) (Client, error)
	DeleteClient(string) error

	CreateAppointment(appt Appointment) (Appointment, error)
	ListServices([]Param) ([]Service, error)
	GetClientServicesCount(string) ([]ServiceSummary, error)

	ListServiceTypes([]Param) ([]ServiceType, error)
	CreateServiceType(st ServiceType) (ServiceType, error)
}
