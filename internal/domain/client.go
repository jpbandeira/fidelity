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
	ClientEntity string = "client"
)

func (c Client) validateClient() error {
	if c.Name == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
			Field:  ferros.NameField,
			Msg:    ferros.EmptyErrorString,
			Entity: ClientEntity,
		})
	}

	if c.Email == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
			Field:  ferros.EmailField,
			Msg:    ferros.EmptyErrorString,
			Entity: ClientEntity,
		})
	}

	if c.Phone == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
			Field:  ferros.PhoneField,
			Msg:    ferros.EmptyErrorString,
			Entity: ClientEntity,
		})
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
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
			Field:  ferros.IdField,
			Msg:    ferros.EmptyErrorString,
			Entity: ClientEntity,
		})
	}

	err := a.db.DeleteClient(id)
	if err != nil {
		return fmt.Errorf(
			ferros.ErrFormatString, ferros.ErrNotFound, &ferros.NotFoundError{
				Entity: ClientEntity,
			},
		)
	}

	return nil
}
