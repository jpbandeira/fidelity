package domain

import (
	"fmt"
	"time"

	ferros "github.com/jp/fidelity/internal/pkg/errors"
)

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

type ClientServiceTypeCount struct {
	ServiceType ServiceType
	Client      Client
	Count       int
}

func (s Service) validateService() error {
	if s.Client.ID == "" {
		return fmt.Errorf("%w: %s", ferros.ErrInvalidParameter, "Missing client reference")
	}

	if s.Attendant.ID == "" {
		return fmt.Errorf("%w: %s", ferros.ErrInvalidParameter, "Missing attendant reference")
	}

	if s.ServiceType == "" {
		return fmt.Errorf("%w: %s", ferros.ErrInvalidParameter, "Missing service type")
	}

	if s.PaymentType == "" {
		return fmt.Errorf("%w: %s", ferros.ErrInvalidParameter, "Missing payment type")
	}

	if s.Description == "" {
		return fmt.Errorf("%w: %s", ferros.ErrInvalidParameter, "Missing service description")
	}

	if s.ServiceDate.String() == "" {
		return fmt.Errorf("%w: %s", ferros.ErrInvalidParameter, "Missing service date")
	}

	if s.Price < 0 {
		return fmt.Errorf("%w: %s", ferros.ErrInvalidParameter, "Service price cannot be negative")
	}

	if s.Price == 0 {
		return fmt.Errorf("%w: %s", ferros.ErrInvalidParameter, "Service price should be greather than 0")
	}

	return nil
}

func (a *actions) CreateService(service Service) (Service, error) {
	err := service.validateService()
	if err != nil {
		return Service{}, err
	}

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

func (a *actions) GetClientServicesCount(cliendUUID string) ([]ClientServiceTypeCount, error) {
	return a.db.GetClientServicesCount(cliendUUID)
}
