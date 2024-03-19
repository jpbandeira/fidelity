package domain

import "context"

type Actions interface {
	CreateClient(context.Context, Client) (Client, error)
}

type actions struct {
	db RepoPersister
}

func ProviderService(db RepoPersister) Actions {
	return &actions{
		db: db,
	}
}
