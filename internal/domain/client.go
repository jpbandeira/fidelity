package domain

import (
	"fmt"
	"time"

	ferros "github.com/jp/fidelity/internal/pkg/errors"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
}

const (
	nameField  string = "name"
	emailField string = "email"
	phoneField string = "phone"

	emptyErrorString string = "Empty"
)

func (c Client) validateClient() error {
	if c.Name == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field: nameField,
			Msg:   emptyErrorString,
		}.Error())
	}

	if c.Email == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field: emailField,
			Msg:   emptyErrorString,
		}.Error())
	}

	if c.Phone == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, ferros.ValidationError{
			Field: emailField,
			Msg:   emptyErrorString,
		}.Error())
	}

	return nil
}

func (a actions) CreateClient(client Client) (Client, error) {
	err := client.validateClient()
	if err != nil {
		return Client{}, err
	}

	c, err := a.db.CreateClient(client)
	if err != nil {
		return Client{}, err
	}

	return c, nil
}

func (a actions) UpdateClient(client Client) (Client, error) {
	err := client.validateClient()
	if err != nil {
		return Client{}, err
	}

	c, err := a.db.UpdateClient(client)
	if err != nil {
		return Client{}, err
	}

	return c, nil
}

func (a actions) ListClients(params []Param) ([]Client, error) {
	return a.db.ListClients(params)
}

func (a actions) DeleteClient(id string) error {
	if id == "" {
		return fmt.Errorf("%w: %s", ferros.ErrInvalidParameter, "Empty client id")
	}

	return a.db.DeleteClient(id)
}
