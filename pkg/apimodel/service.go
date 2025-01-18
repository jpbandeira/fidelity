package apimodel

import "time"

type Service struct {
	ID          string    `json:"ID"`
	Client      Client    `json:"Client"`
	Attendant   Attendant `json:"Attendant"`
	Price       float32   `json:"Price"`
	ServiceType string    `json:"ServiceType"`
	PaymentType string    `json:"PaymentType"`
	Description string    `json:"Description"`
	ServiceDate time.Time `json:"ServiceDate"`
}
