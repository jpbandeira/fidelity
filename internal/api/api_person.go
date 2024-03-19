package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/pkg/apimodel"
)

func apiToDomain(c apimodel.Person) domain.Person {
	return domain.Person{
		Name:  c.Name,
		Email: c.Email,
		Phone: c.Phone,
	}
}

func domainToAPI(c domain.Person) apimodel.Person {
	return apimodel.Person{
		ID:    c.ID,
		Name:  c.Name,
		Email: c.Email,
		Phone: c.Phone,
	}
}

// createPerson - Create a person
func (h *handler) createPerson(c *gin.Context) {
	var personAPI apimodel.Person

	ctx := context.Background()

	err := c.BindJSON(&personAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	person, err := h.actions.CreatePerson(ctx, apiToDomain(personAPI))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, domainToAPI(person))
}
