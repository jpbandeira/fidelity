package domain

import (
	"errors"
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

const (
	ServiceEntity string = "service"
)

func (s Service) validateService() error {
	if s.Client.ID == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field:  ferros.IdField,
			Msg:    ferros.EmptyErrorString,
			Entity: ClientEntity,
		})
	}

	if s.Attendant.ID == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field:  ferros.IdField,
			Msg:    ferros.EmptyErrorString,
			Entity: AttendantEntity,
		})
	}

	if s.ServiceType == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field:  ferros.TypeField,
			Msg:    ferros.EmptyErrorString,
			Entity: ServiceEntity,
		})
	}

	if s.PaymentType == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field:  ferros.PaymentTypeField,
			Msg:    ferros.EmptyErrorString,
			Entity: ServiceEntity,
		})
	}

	if s.Description == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field:  ferros.DescriptionField,
			Msg:    ferros.EmptyErrorString,
			Entity: ServiceEntity,
		})
	}

	if s.ServiceDate.String() == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field:  ferros.DateField,
			Msg:    ferros.EmptyErrorString,
			Entity: ServiceEntity,
		})
	}

	if s.Price < 0 {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field:  ferros.PriceField,
			Msg:    ferros.CannotBeNegativeErrorString,
			Entity: ServiceEntity,
		})
	}

	if s.Price == 0 {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field:  ferros.PriceField,
			Msg:    ferros.ShouldBeGreaterThanErrorString,
			Entity: ServiceEntity,
		})
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
		if errors.Is(err, ferros.ErrNotFound) {
			return Service{}, fmt.Errorf(
				ferros.ErrFormatString, ferros.ErrNotFound, ferros.NotFoundError{
					Entity: ClientEntity,
				},
			)
		}

		return Service{}, err
	}

	attendant, err := a.db.GetAttendant(service.Attendant.ID)
	if err != nil {
		if errors.Is(err, ferros.ErrNotFound) {
			return Service{}, fmt.Errorf(
				ferros.ErrFormatString, ferros.ErrNotFound, ferros.NotFoundError{
					Entity: AttendantEntity,
				},
			)
		}

		return Service{}, err
	}

	service.Client = client
	service.Attendant = attendant

	return a.db.CreateService(service)
}

func (a *actions) ListServicesByClient(clientID string, params []Param) ([]Service, error) {
	if clientID == "" {
		return []Service{},
			fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
				Field:  ferros.IdField,
				Msg:    ferros.EmptyErrorString,
				Entity: ClientEntity,
			})
	}

	return a.db.ListServicesByClient(clientID, params)
}

func (a *actions) GetClientServicesCount(clientID string) ([]ClientServiceTypeCount, error) {
	if clientID == "" {
		return []ClientServiceTypeCount{},
			fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
				Field:  ferros.IdField,
				Msg:    ferros.EmptyErrorString,
				Entity: ClientEntity,
			})
	}

	return a.db.GetClientServicesCount(clientID)
}
