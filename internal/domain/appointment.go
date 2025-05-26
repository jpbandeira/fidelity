package domain

import (
	"fmt"
	"time"

	ferros "github.com/jp/fidelity/internal/pkg/errors"
)

type Appointment struct {
	ID        string
	Client    Client
	Attendant Attendant
	Services  []Service
}

type Service struct {
	ID          string
	Name        string
	Price       float32
	PaymentType string
	Description string
	ServiceDate time.Time
	Client      Client
	Attendant   Attendant
}

type ClientServiceTypeCount struct {
	ServiceType ServiceType
	Client      Client
	Count       int
}

func (a *actions) ListServices(params []Param) ([]Service, error) {
	return a.db.ListServices(params)
}

func (a *actions) GetClientServicesCount(clientID string) ([]ClientServiceTypeCount, error) {
	if clientID == "" {
		return []ClientServiceTypeCount{},
			fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
				Field:  ferros.IdField,
				Msg:    ferros.EmptyErrorString,
				Entity: ClientEntity,
			})
	}

	return a.db.GetClientServicesCount(clientID)
}
