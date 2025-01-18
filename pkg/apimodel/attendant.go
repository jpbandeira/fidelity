package apimodel

import "time"

type Attendant struct {
	ID        string    `json:"ID"`
	Name      string    `json:"Name"`
	Email     string    `json:"Email"`
	Phone     string    `json:"Phone"`
	CreatedAt time.Time `json:"CreatedAt"`
}
