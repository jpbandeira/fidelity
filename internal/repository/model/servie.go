package model

import "gorm.io/gorm"

type Service struct {
	*gorm.Model
	// UUID is a unique globally id
	UUID string `gorm:"unique;not null;index"`

	// ClientID  represents the relationship between the client user and service
	ClientID uint `gorm:"not null; index:idx_user_client_service;"`
	Client   User `gorm:"references:id;"`

	// AttedantID  represents the relationship between the attendant user and service
	AttedantID uint `gorm:"not null; index:idx_user_attendant_service;"`
	Attendant  User `gorm:"references:id;"`

	Price       float32 `gorm:"not null"`
	ServiceType uint    `gorm:"not null"`
	PaymentType uint    `gorm:"not null"`
	Description string
}
