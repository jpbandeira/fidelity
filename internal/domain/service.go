package domain

import "time"

type Service struct {
	ID          string
	Client      Client
	Attendant   Attendant
	Price       float32
	ServiceType string
	PaymentType string
	Description string
	ServiceDate time.Time
}

type ClientServiceCount struct {
	ServiceType  ServiceType
	Client       Client
	ServiceCount int
}

func (a *actions) CreateService(service Service) (Service, error) {
	client, err := a.db.GetClient(service.Client.ID)
	if err != nil {
		return Service{}, err
	}

	attendant, err := a.db.GetAttendant(service.Attendant.ID)
	if err != nil {
		return Service{}, err
	}

	service.Client = client
	service.Attendant = attendant

	return a.db.CreateService(service)
}

func (a *actions) ListServicesByClient(clientID string, params []Param) ([]Service, error) {
	return a.db.ListServicesByClient(clientID, params)
}

func (a *actions) GetClientServicesCount(cliendUUID string) ([]ClientServiceCount, error) {
	return a.db.GetClientServicesCount(cliendUUID)
}
