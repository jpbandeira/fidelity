package model

import (
	"github.com/jp/fidelity/internal/domain"
	"gorm.io/gorm"
)

type Attendant struct {
	gorm.Model
	// UUID is a unique globally id
	UUID  string `gorm:"unique;not null;index"`
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
	Phone string `gorm:"unique;not null"`
}

func (p Attendant) RepoToDomain() domain.Attendant {
	return domain.Attendant{
		ID:        p.UUID,
		Name:      p.Name,
		Email:     p.Email,
		Phone:     p.Phone,
		CreatedAt: p.CreatedAt,
	}
}

func AttendantDomainToRepo(p domain.Attendant) Attendant {
	return Attendant{
		Model: gorm.Model{
			CreatedAt: p.CreatedAt,
		},
		UUID:  p.ID,
		Name:  p.Name,
		Email: p.Email,
		Phone: p.Phone,
	}
}
