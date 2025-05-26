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

type ServiceSummary struct {
	ServiceType ServiceType
	Client      Client
	Count       int
	TotalPrice  float32
}

const (
	AppointmentEntity string = "appointment"
)

func validateAppointment(appointment Appointment) error {
	if appointment.Client.ID == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
			Field:  ferros.IdField,
			Msg:    ferros.EmptyErrorString,
			Entity: ClientEntity,
		})
	}

	if appointment.Attendant.ID == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
			Field:  ferros.IdField,
			Msg:    ferros.EmptyErrorString,
			Entity: AttendantEntity,
		})
	}

	for _, a := range appointment.Services {
		if a.Name == "" {
			return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
				Field:  ferros.NameField,
				Msg:    ferros.EmptyErrorString,
				Entity: AppointmentEntity,
			})
		}

		if a.PaymentType == "" {
			return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
				Field:  ferros.PaymentTypeField,
				Msg:    ferros.EmptyErrorString,
				Entity: AppointmentEntity,
			})
		}

		if a.ServiceDate.String() == "" {
			return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
				Field:  ferros.DateField,
				Msg:    ferros.EmptyErrorString,
				Entity: AppointmentEntity,
			})
		}

		if a.Price < 0 {
			return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
				Field:  ferros.PriceField,
				Msg:    ferros.CannotBeNegativeErrorString,
				Entity: AppointmentEntity,
			})
		}

		if a.Price == 0 {
			return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
				Field:  ferros.PriceField,
				Msg:    ferros.ShouldBeGreaterThanErrorString,
				Entity: AppointmentEntity,
			})
		}
	}
	return nil
}

func (a *actions) CreateAppointment(appt Appointment) (Appointment, error) {
	if err := validateAppointment(appt); err != nil {
		return Appointment{}, err
	}

	client, err := a.db.GetClient(appt.Client.ID)
	if err != nil {
		return Appointment{}, err
	}

	attendant, err := a.db.GetAttendant(appt.Attendant.ID)
	if err != nil {
		return Appointment{}, err
	}

	appt.Client = client
	appt.Attendant = attendant
	return a.db.CreateAppointment(appt)
}

func (a *actions) ListServices(params []Param) ([]Service, error) {
	return a.db.ListServices(params)
}

func (a *actions) GetServiceSummary(clientID string) ([]ServiceSummary, error) {
	if clientID == "" {
		return []ServiceSummary{},
			fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
				Field:  ferros.IdField,
				Msg:    ferros.EmptyErrorString,
				Entity: ClientEntity,
			})
	}

	return a.db.GetClientServicesCount(clientID)
}
