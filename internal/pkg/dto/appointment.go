package dto

import "time"

type Appointment struct {
	ID          string    `json:"id"`
	Client      Client    `json:"client" binding:"required"`
	AttendantID string    `json:"attendantID" binding:"required"`
	Services    []Service `json:"services" binding:"required"`
}

type Service struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Price       float32   `json:"price" binding:"required"`
	PaymentType string    `json:"paymentType" binding:"required"`
	Description string    `json:"description"`
	ServiceDate time.Time `json:"serviceDate" binding:"required"`
	ClientID    string    `json:"clientID,omitempty"`
	ClientName  string    `json:"clientName,omitempty"`
	AttendantID string    `json:"attendantID,omitempty"`
}

type ServiceSummary struct {
	Name       string  `json:"name"`
	Count      int     `json:"count"`
	TotalPrice float32 `json:"totalPrice"`
}

type ServiceListResponse struct {
	Services         []Service        `json:"services"`
	ServiceSummaries []ServiceSummary `json:"serviceSummaries"`
}
