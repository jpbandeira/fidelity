package domain

// import (
// 	"errors"
// 	"fmt"
// 	"time"

// 	ferros "github.com/jp/fidelity/internal/pkg/errors"
// )

// type Service struct {
// 	ID          string
// 	Price       float32
// 	ServiceType string
// 	PaymentType string
// 	Description string
// 	ServiceDate time.Time
// }

// type ServiceBatch struct {
// 	Client    Client
// 	Attendant Attendant
// 	Items     []Service
// }

// type ClientServiceTypeCount struct {
// 	ServiceType ServiceType
// 	Client      Client
// 	Count       int
// }

// const (
// 	ServiceEntity string = "service"
// )

// func validateServiceBatch(serviceBatch ServiceBatch) error {
// 	if serviceBatch.Client.ID == "" {
// 		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
// 			Field:  ferros.IdField,
// 			Msg:    ferros.EmptyErrorString,
// 			Entity: ClientEntity,
// 		})
// 	}

// 	if serviceBatch.Attendant.ID == "" {
// 		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
// 			Field:  ferros.IdField,
// 			Msg:    ferros.EmptyErrorString,
// 			Entity: AttendantEntity,
// 		})
// 	}

// 	for _, s := range serviceBatch.Items {
// 		if s.ServiceType == "" {
// 			return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
// 				Field:  ferros.TypeField,
// 				Msg:    ferros.EmptyErrorString,
// 				Entity: ServiceEntity,
// 			})
// 		}

// 		if s.PaymentType == "" {
// 			return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
// 				Field:  ferros.PaymentTypeField,
// 				Msg:    ferros.EmptyErrorString,
// 				Entity: ServiceEntity,
// 			})
// 		}

// 		if s.ServiceDate.String() == "" {
// 			return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
// 				Field:  ferros.DateField,
// 				Msg:    ferros.EmptyErrorString,
// 				Entity: ServiceEntity,
// 			})
// 		}

// 		if s.Price < 0 {
// 			return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
// 				Field:  ferros.PriceField,
// 				Msg:    ferros.CannotBeNegativeErrorString,
// 				Entity: ServiceEntity,
// 			})
// 		}

// 		if s.Price == 0 {
// 			return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
// 				Field:  ferros.PriceField,
// 				Msg:    ferros.ShouldBeGreaterThanErrorString,
// 				Entity: ServiceEntity,
// 			})
// 		}
// 	}
// 	return nil
// }

// func (a *actions) CreateServiceBatch(servicesBatch ServiceBatch) (ServiceBatch, error) {
// 	err := validateServiceBatch(servicesBatch)
// 	if err != nil {
// 		return ServiceBatch{}, err
// 	}

// 	client, err := a.db.GetClient(servicesBatch.Client.ID)
// 	if err != nil {
// 		if errors.Is(err, ferros.ErrNotFound) {
// 			return ServiceBatch{}, fmt.Errorf(
// 				ferros.ErrFormatString, ferros.ErrNotFound, &ferros.NotFoundError{
// 					Entity: ClientEntity,
// 				},
// 			)
// 		}

// 		return ServiceBatch{}, err
// 	}

// 	attendant, err := a.db.GetAttendant(servicesBatch.Attendant.ID)
// 	if err != nil {
// 		if errors.Is(err, ferros.ErrNotFound) {
// 			return ServiceBatch{}, fmt.Errorf(
// 				ferros.ErrFormatString, ferros.ErrNotFound, &ferros.NotFoundError{
// 					Entity: AttendantEntity,
// 				},
// 			)
// 		}

// 		return ServiceBatch{}, err
// 	}

// 	servicesBatch.Client = client
// 	servicesBatch.Attendant = attendant

// 	return a.db.CreateServiceBatch(servicesBatch)
// }

// func (a *actions) ListServices(params []Param) (ServiceBatch, error) {
// 	return a.db.ListServices(params)
// }

// func (a *actions) GetClientServicesCount(clientID string) ([]ClientServiceTypeCount, error) {
// 	if clientID == "" {
// 		return []ClientServiceTypeCount{},
// 			fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
// 				Field:  ferros.IdField,
// 				Msg:    ferros.EmptyErrorString,
// 				Entity: ClientEntity,
// 			})
// 	}

// 	return a.db.GetClientServicesCount(clientID)
// }
