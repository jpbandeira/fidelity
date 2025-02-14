package apimodel

import "time"

type Service struct {
	ID          string    `json:"id"`
	Client      Client    `json:"client"`
	Attendant   Attendant `json:"attendant"`
	Price       float32   `json:"price"`
	ServiceType string    `json:"serviceType"`
	PaymentType string    `json:"paymentType"`
	Description string    `json:"description"`
	ServiceDate time.Time `json:"serviceDate"`
}

type ServiceTypeCount struct {
	ServiceType string `json:"serviceType"`
	Count       int    `json:"count"`
}

type ServiceList struct {
	Items             []Service          `json:"items"`
	ServiceTypesCount []ServiceTypeCount `json:"serviceTypes"`
}
