package model

import (
	"time"

	"github.com/jp/fidelity/internal/domain"
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	UUID string `gorm:"unique;not null;index"`

	ClientUUID string `gorm:"not null;index"`
	Client     Client `gorm:"foreignKey:ClientUUID;references:UUID;constraint:OnDelete:CASCADE;"`

	AttendantUUID string    `gorm:"not null;index"`
	Attendant     Attendant `gorm:"foreignKey:AttendantUUID;references:UUID;constraint:OnDelete:CASCADE;"`

	Services []Service `gorm:"foreignKey:AppointmentUUID;references:UUID;constraint:OnDelete:CASCADE;"`
}

type Service struct {
	gorm.Model
	UUID string `gorm:"unique;not null;index"`

	// Relacionamento com o Atendimento
	AppointmentUUID string      `gorm:"not null;index"`
	Appointment     Appointment `gorm:"foreignKey:AppointmentUUID;references:UUID;constraint:OnDelete:CASCADE;"`

	// Tipo de serviço
	ServiceTypeID uint        `gorm:"not null;index"`
	ServiceType   ServiceType `gorm:"foreignKey:ServiceTypeID;references:ID;constraint:OnUpdate:CASCADE;"`

	// Dados do serviço
	Price       float32            `gorm:"not null"`
	PaymentType domain.PaymentType `gorm:"not null"`
	Description string
	ServiceDate time.Time `gorm:"not null"`
}

type ServiceSummary struct {
	ServiceTypeID uint        `gorm:"not null;index:idx_service_type_service_count;uniqueIndex:idx_service_type_client;"`
	ServiceType   ServiceType `gorm:"foreignKey:ServiceTypeID;references:ID;constraint:OnUpdate:CASCADE;"`

	ClientUUID string `gorm:"not null;index:idx_user_client_service_count;uniqueIndex:idx_service_type_client;"`
	Client     Client `gorm:"foreignKey:ClientUUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Count      int     `gorm:"not null"`
	TotalPrice float32 `gorm:"not null"`
}

func ServiceRepoToDomain(services []Service) []domain.Service {
	var serviceList = make([]domain.Service, 0, len(services))
	for _, s := range services {
		serviceList = append(serviceList, domain.Service{
			ID:          s.UUID,
			Name:        s.ServiceType.Name,
			Price:       s.Price,
			PaymentType: s.PaymentType.String(),
			Description: s.Description,
			ServiceDate: s.ServiceDate,
			Client:      s.Appointment.Client.RepoToDomain(),
			Attendant:   s.Appointment.Attendant.RepoToDomain(),
		})
	}
	return serviceList
}

func (csc ServiceSummary) RepoToDomain() domain.ClientServiceTypeCount {
	return domain.ClientServiceTypeCount{
		ServiceType: domain.ServiceType{
			Name: csc.ServiceType.Name,
		},
		Client: csc.Client.RepoToDomain(),
		Count:  csc.Count,
	}
}

func (a Appointment) RepoToDomain() domain.Appointment {
	services := make([]domain.Service, 0, len(a.Services))
	for _, s := range a.Services {
		services = append(services, domain.Service{
			ID:          s.UUID,
			Name:        s.ServiceType.Name,
			Price:       s.Price,
			PaymentType: s.PaymentType.String(),
			Description: s.Description,
			ServiceDate: s.ServiceDate,
			Client:      a.Client.RepoToDomain(),
			Attendant:   a.Attendant.RepoToDomain(),
		})
	}

	return domain.Appointment{
		ID:        a.UUID,
		Client:    a.Client.RepoToDomain(),
		Attendant: a.Attendant.RepoToDomain(),
		Services:  services,
	}
}
