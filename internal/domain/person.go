package domain

import (
	"context"
)

type Person struct {
	ID    string
	Name  string
	Email string
	Phone string
}

func (a actions) CreatePerson(ctx context.Context, person Person) (Person, error) {
	person, err := a.db.CreatePerson(ctx, person)
	if err != nil {
		return Person{}, err
	}

	return person, nil
}
