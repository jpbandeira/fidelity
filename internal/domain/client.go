package domain

import "time"

type Client struct {
	ID        string
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
}

func (a actions) CreateClient(client Client) (Client, error) {
	client, err := a.db.CreateClient(client)
	if err != nil {
		return Client{}, err
	}

	return client, nil
}

func (a actions) UpdateClient(client Client) (Client, error) {
	client, err := a.db.UpdateClient(client)
	if err != nil {
		return Client{}, err
	}

	return client, nil
}

func (a actions) ListClients(params []Param) ([]Client, error) {
	return a.db.ListClients(params)
}

func (a actions) DeleteClient(id string) error {
	return a.db.DeleteClient(id)
}
