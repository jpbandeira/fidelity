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
