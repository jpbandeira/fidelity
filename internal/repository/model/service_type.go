package model

import "gorm.io/gorm"

type ServiceType struct {
	*gorm.Model
	Description string `gorm:"not null"`
}
