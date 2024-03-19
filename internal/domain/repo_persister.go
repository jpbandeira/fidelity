package domain

import "context"

type RepoPersister interface {
	CreatePerson(context.Context, Person) (Person, error)
}
