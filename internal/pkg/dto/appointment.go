package dto

import "time"

type Appointment struct {
	ID        string    `json:"id"`
	Client    Client    `json:"client" binding:"required"`
	Attendant Attendant `json:"attendant" binding:"required"`
	Services  []Service `json:"services" binding:"required"`
}

type Service struct {
	ID            string    `json:"id"`
	Name          string    `json:"name" binding:"required"`
	Price         float32   `json:"price" binding:"required"`
	PaymentType   string    `json:"paymentType" binding:"required"`
	Description   string    `json:"description"`
	ServiceDate   time.Time `json:"serviceDate" binding:"required"`
	ClientID      string    `json:"clientID,omitempty"`
	ClientName    string    `json:"clientName,omitempty"`
	AttendantID   string    `json:"attendantID,omitempty"`
	AttendantName string    `json:"attendantName,omitempty"`
}

type ServiceSummary struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type ServiceListResponse struct {
	Services         []Service        `json:"services"`
	ServiceSummaries []ServiceSummary `json:"serviceSummaries"`
}
