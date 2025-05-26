package domain

type Repository interface {
	CreateClient(Client) (Client, error)
	UpdateClient(Client) (Client, error)
	ListClients([]Param) ([]Client, error)
	GetClient(string) (Client, error)
	DeleteClient(string) error

	CreateAttendant(Attendant) (Attendant, error)
	UpdateAttendant(Attendant) (Attendant, error)
	ListAttendants([]Param) ([]Attendant, error)
	GetAttendant(string) (Attendant, error)
	DeleteAttendant(string) error

	// CreateServiceBatch(ServiceBatch) (ServiceBatch, error)
	ListServices([]Param) ([]Service, error)
	GetClientServicesCount(string) ([]ClientServiceTypeCount, error)

	ListServiceTypes([]Param) ([]ServiceType, error)
	CreateServiceType(st ServiceType) (ServiceType, error)
}
