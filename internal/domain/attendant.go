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

const (
	AttendantEntity string = "attendant"
)

func (a Attendant) validateClient() error {
	if a.Name == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field:  ferros.NameField,
			Msg:    ferros.EmptyErrorString,
			Entity: AttendantEntity,
		}.Error())
	}

	if a.Email == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field:  ferros.EmailField,
			Msg:    ferros.EmptyErrorString,
			Entity: AttendantEntity,
		}.Error())
	}

	if a.Phone == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field:  ferros.PhoneField,
			Msg:    ferros.EmptyErrorString,
			Entity: AttendantEntity,
		}.Error())
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
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field:  ferros.IdField,
			Msg:    ferros.EmptyErrorString,
			Entity: AttendantEntity,
		}.Error())
	}

	err := a.db.DeleteAttendant(id)
	if err != nil {
		return fmt.Errorf(
			ferros.ErrFormatString, ferros.ErrNotFound, ferros.NotFoundError{
				Entity: AttendantEntity,
			}.Error(),
		)
	}

	return nil
}
