package repository

import (
	"context"

	"github.com/google/uuid"
	
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/internal/repository/models"
)

func (db *GormRepository) CreatePerson(ctx context.Context, person domain.Person) (domain.Person, error) {
	personModel := models.Person{
		UUID:  uuid.NewString(),
		Name:  person.Name,
		Email: person.Email,
		Phone: person.Phone,
	}

	err := db.Create(&personModel).Error
	if err != nil {
		return domain.Person{}, err
	}

	return personModel.RepoToDomain(), nil
}
