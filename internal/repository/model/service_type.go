package model

import (
	"github.com/jp/fidelity/internal/domain"
	"gorm.io/gorm"
)

type ServiceType struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func (s ServiceType) RepoToDomain() domain.ServiceType {
	return domain.ServiceType{
		Name: s.Name,
	}
}
