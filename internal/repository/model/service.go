package model

// import (
// 	"time"

// 	"github.com/jp/fidelity/internal/domain"
// 	"gorm.io/gorm"
// )

// type Service struct {
// 	gorm.Model
// 	// UUID is a unique globally id
// 	UUID string `gorm:"unique;not null;index"`

// 	// ClientID  represents the relationship between the client user and service
// 	ClientUUID string `gorm:"not null; index:idx_user_client_service;"`
// 	Client     Client `gorm:"foreignKey:ClientUUID;references:UUID;constraint:OnDelete:CASCADE;"`

// 	// AttedantID  represents the relationship between the attendant user and service
// 	AttedantUUID string    `gorm:"not null; index:idx_user_attendant_service;"`
// 	Attendant    Attendant `gorm:"foreignKey:AttedantUUID;references:UUID;constraint:OnDelete:CASCADE;"`

// 	ServiceTypeID uint        `gorm:"not null; index:idx_service_type_service;"`
// 	ServiceType   ServiceType `gorm:"foreignKey:ServiceTypeID;references:ID;constraint:OnUpdate:CASCADE"`

// 	Price       float32            `gorm:"not null"`
// 	PaymentType domain.PaymentType `gorm:"not null"`
// 	Description string
// 	ServiceDate time.Time `gorm:"not null"`
// }

// type ClientServiceTypeCount struct {
// 	// ServiceTypeID  represents the relationship between service type and service count
// 	ServiceTypeID uint        `gorm:"not null; index:idx_service_type_service_count; uniqueIndex:idx_service_type_client;"`
// 	ServiceType   ServiceType `gorm:"foreignKey:ServiceTypeID;references:ID;constraint:OnUpdate:CASCADE"`

// 	// ClientUUID  represents the relationship between client user and service count
// 	ClientUUID string `gorm:"not null; index:idx_user_client_service_count; uniqueIndex:idx_service_type_client;"`
// 	Client     Client `gorm:"foreignKey:ClientUUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

// 	// ServiceCount is the number of service done by client in a specific service type
// 	ServiceCount int `gorm:"not null"`
// }

// func ServiceRepoToDomain(services []Service) []domain.Service {
// 	var serviceList = make([]domain.Service, 0)
// 	for _, s := range services {
// 		serviceList = append(serviceList, domain.Service{
// 			ID:          s.UUID,
// 			Price:       s.Price,
// 			ServiceType: s.ServiceType.Description,
// 			PaymentType: s.PaymentType.String(),
// 			Description: s.Description,
// 			ServiceDate: s.ServiceDate,
// 		})
// 	}

// 	return serviceList
// }

// func (csc ClientServiceTypeCount) RepoToDomain() domain.ClientServiceTypeCount {
// 	return domain.ClientServiceTypeCount{
// 		ServiceType: csc.ServiceType.RepoToDomain(),
// 		Client:      csc.Client.RepoToDomain(),
// 		Count:       csc.ServiceCount,
// 	}
// }
