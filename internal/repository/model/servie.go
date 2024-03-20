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
	ClientID uint `gorm:"not null; index:idx_user_client_service;"`
	Client   User `gorm:"foreignKey:ClientID;references:ID;constraint:OnDelete:CASCADE;"`

	// AttedantID  represents the relationship between the attendant user and service
	AttedantID uint `gorm:"not null; index:idx_user_attendant_service;"`
	Attendant  User `gorm:"foreignKey:AttedantID;references:ID;constraint:OnDelete:CASCADE;"`

	Price       float32 `gorm:"not null"`
	ServiceType uint    `gorm:"not null"`
	PaymentType uint    `gorm:"not null"`
	Description string
}

type ClientServiceCount struct {
	ServiceType uint `gorm:"not null; uniqueIndex:idx_service_type_client;"`
	// ClientID  represents the relationship between the client user and service
	ClientID uint `gorm:"not null; index:idx_user_client_service_count; uniqueIndex:idx_service_type_client;"`
	Client   User `gorm:"foreignKey:ClientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
