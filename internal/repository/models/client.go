package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	// UUID is a unique globally id
	UUID  string `gorm:"unique;not null;index"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
	Phone string `gorm:"unique;not null"`
}
