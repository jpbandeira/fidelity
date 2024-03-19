package domain

import (
	"context"
)

type Client struct {
	ID    string
	Name  string
	Email string
	Phone string
}

func (a actions) CreateClient(ctx context.Context, client Client) (Client, error) {
	return Client{}, nil
}
