package domain

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound            = errors.New("not found")
	ErrInvalidParameter    = errors.New("invalid parameter")
	ErrInternalServerError = errors.New("internal server error")
)

const (
	ErrFormatString string = "%w: %w"

	NameField  string = "name"
	EmailField string = "email"
	PhoneField string = "phone"
	IdField    string = "id"

	TypeField        string = "type"
	PaymentTypeField string = "payment type"
	DescriptionField string = "description"
	DateField        string = "date"
	PriceField       string = "price"

	EmptyErrorString               string = "empty"
	MissingErrorString             string = "missing"
	CannotBeNegativeErrorString    string = "cannot be negative"
	ShouldBeGreaterThanErrorString string = "should be greater than 0"
)

type ValidationError struct {
	Field  string
	Msg    string
	Entity string
}

func (vErr ValidationError) Error() string {
	return fmt.Sprintf("%s %s: %s", vErr.Entity, vErr.Field, vErr.Msg)
}

type NotFoundError struct {
	Entity string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Entity)
}
