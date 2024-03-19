package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/pkg/apimodel"
)

func domainToAPI(client domain.Client) apimodel.Client {
	return apimodel.Client{
		ID:    client.ID,
		Name:  client.Name,
		Email: client.Email,
		Phone: client.Phone,
	}
}

// createScopeGroup - Create a scope group
func (h *handler) createClient(c *gin.Context) {
	var clientAPI apimodel.Client

	ctx := context.Background()

	err := c.BindJSON(&clientAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client, err := h.actions.CreateClient(ctx, domain.Client{})
	if err != nil {
		c.JSON(1, err)
	}

	c.JSON(http.StatusCreated, domainToAPI(client))
}
