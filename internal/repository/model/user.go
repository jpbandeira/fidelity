package model

import (
	"github.com/jp/fidelity/internal/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// UUID is a unique globally id
	UUID  string `gorm:"unique;not null;index"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
	Phone string `gorm:"unique;not null"`
}

func (p User) RepoToDomain() domain.User {
	return domain.User{
		ID:    p.UUID,
		Name:  p.Name,
		Email: p.Email,
		Phone: p.Phone,
	}
}

func UserDomainToRepo(p domain.User) User {
	return User{
		UUID:  p.ID,
		Name:  p.Name,
		Email: p.Email,
		Phone: p.Phone,
	}
}
