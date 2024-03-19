package models

import (
	"github.com/jp/fidelity/internal/domain"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	// UUID is a unique globally id
	UUID  string `gorm:"unique;not null;index"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
	Phone string `gorm:"unique;not null"`
}

func (p Person) RepoToDomain() domain.Person {
	return domain.Person{
		ID:    p.UUID,
		Name:  p.Name,
		Email: p.Email,
		Phone: p.Phone,
	}
}
