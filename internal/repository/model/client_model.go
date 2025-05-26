package model

import (
	"github.com/jp/fidelity/internal/domain"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	// UUID is a unique globally id
	UUID  string `gorm:"unique;not null;index"`
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
	Phone string `gorm:"unique;not null"`
}

func (p Client) RepoToDomain() domain.Client {
	return domain.Client{
		ID:        p.UUID,
		Name:      p.Name,
		Email:     p.Email,
		Phone:     p.Phone,
		CreatedAt: p.CreatedAt,
	}
}

func ClientDomainToRepo(p domain.Client) Client {
	return Client{
		Model: gorm.Model{
			CreatedAt: p.CreatedAt,
		},
		UUID:  p.ID,
		Name:  p.Name,
		Email: p.Email,
		Phone: p.Phone,
	}
}
