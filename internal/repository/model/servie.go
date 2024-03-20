package model

import (
	"github.com/jp/fidelity/internal/domain"
	"gorm.io/gorm"
)

type Service struct {
	*gorm.Model
	// UUID is a unique globally id
	UUID string `gorm:"unique;not null;index"`

	// ClientID  represents the relationship between the client user and service
	ClientID uint8 `gorm:"not null; index:idx_user_client_service;"`
	Client   User  `gorm:"references:id;"`

	// AttedantID  represents the relationship between the attendant user and service
	AttedantID uint8 `gorm:"not null; index:idx_user_attendant_service;"`
	Attendant  User  `gorm:"references:id;"`

	Price       float32 `gorm:"not null"`
	ServiceType uint8   `gorm:"not null"`
	PaymentType uint8   `gorm:"not null"`
	Description string
}

type ClientServiceCount struct {
	ServiceType uint8 `gorm:"not null"`
	// ClientID  represents the relationship between the client user and service
	ClientID uint8 `gorm:"not null; index:idx_user_client_service_count;"`
	Client   User  `gorm:"references:id;"`
	// ServiceCount is the number of service done by client in a specific service type
	ServiceCount int `gorm:"not null"`
}

func (s Service) RepoToDomain() domain.Service {
	return domain.Service{
		ID:          s.UUID,
		Client:      s.Client.RepoToDomain(),
		Attendant:   s.Attendant.RepoToDomain(),
		Price:       s.Price,
		ServiceType: s.ServiceType,
		PaymentType: s.PaymentType,
		Description: s.Description,
	}
}
