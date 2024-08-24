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

	ServiceTypeID uint        `gorm:"not null; index:idx_service_type_service;"`
	ServiceType   ServiceType `gorm:"foreignKey:ServiceTypeID;references:ID;constraint:OnUpdate:CASCADE"`

	Price       float32            `gorm:"not null"`
	PaymentType domain.PaymentType `gorm:"not null"`
	Description string
}

type ClientServiceCount struct {
	// ServiceTypeID  represents the relationship between service type and service count
	ServiceTypeID uint        `gorm:"not null; index:idx_service_type_service_count; uniqueIndex:idx_service_type_client;"`
	ServiceType   ServiceType `gorm:"foreignKey:ServiceTypeID;references:ID;constraint:OnUpdate:CASCADE"`

	// ClientID  represents the relationship between client user and service count
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
		ServiceType: s.ServiceType.Description,
		PaymentType: s.PaymentType.String(),
		Description: s.Description,
	}
}
