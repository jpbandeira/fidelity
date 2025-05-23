package domain

import (
	"fmt"
	"time"

	ferros "github.com/jp/fidelity/internal/pkg/errors"
)

type Attendant struct {
	ID        string
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
}

func (a Attendant) validateClient() error {
	if a.Name == "" {
		return fmt.Errorf("%w: %s", ferros.ErrInvalidParameter, "Empty attendant name")
	}

	if a.Email == "" {
		return fmt.Errorf("%w: %s", ferros.ErrInvalidParameter, "Empty attendant email")
	}

	if a.Phone == "" {
		return fmt.Errorf("%w: %s", ferros.ErrInvalidParameter, "Empty attendant phone")
	}

	return nil
}

func (a actions) CreateAttendant(attendant Attendant) (Attendant, error) {
	err := attendant.validateClient()
	if err != nil {
		return Attendant{}, err
	}

	att, err := a.db.CreateAttendant(attendant)
	if err != nil {
		return Attendant{}, err
	}

	return att, nil
}

func (a actions) UpdateAttendant(attendant Attendant) (Attendant, error) {
	err := attendant.validateClient()
	if err != nil {
		return Attendant{}, err
	}

	att, err := a.db.UpdateAttendant(attendant)
	if err != nil {
		return Attendant{}, err
	}

	return att, nil
}

func (a actions) ListAttendants(params []Param) ([]Attendant, error) {
	return a.db.ListAttendants(params)
}

func (a actions) DeleteAttendant(id string) error {
	if id == "" {
		return fmt.Errorf("%w: %s", ferros.ErrInvalidParameter, "Empty attendant id")
	}

	return a.db.DeleteAttendant(id)
}
