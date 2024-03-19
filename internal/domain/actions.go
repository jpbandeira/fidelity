package domain

import "context"

type Actions interface {
	CreatePerson(context.Context, Person) (Person, error)
}

type actions struct {
	db RepoPersister
}

func ProviderService(db RepoPersister) Actions {
	return &actions{
		db: db,
	}
}
