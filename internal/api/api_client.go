package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/pkg/apimodel"
)

func clientAPIToDomain(c apimodel.Client) domain.Client {
	return domain.Client{
		ID:        c.ID,
		Name:      c.Name,
		Email:     c.Email,
		Phone:     c.Phone,
		CreatedAt: c.CreatedAt,
	}
}

func clientDomainToAPI(c domain.Client) apimodel.Client {
	return apimodel.Client{
		ID:        c.ID,
		Name:      c.Name,
		Email:     c.Email,
		Phone:     c.Phone,
		CreatedAt: c.CreatedAt,
	}
}

// createClient - Create a Client
func (h *handler) createClient(c *gin.Context) {
	var clientAPI apimodel.Client

	err := c.BindJSON(&clientAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client, err := h.actions.CreateClient(clientAPIToDomain(clientAPI))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, clientDomainToAPI(client))
}

// updateClient - Update a Client
func (h *handler) updateClient(c *gin.Context) {
	var clientAPI apimodel.Client

	err := c.BindJSON(&clientAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client, err := h.actions.UpdateClient(clientAPIToDomain(clientAPI))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, clientDomainToAPI(client))
}

// listClient - List a list of Client
func (h *handler) listClient(c *gin.Context) {
	qps := []domain.Param{}
	for key, value := range c.Request.URL.Query() {
		for _, v := range value {
			qps = append(qps, domain.Param{Key: key, Value: v})
		}
	}

	clients, err := h.actions.ListClients(qps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, clients)
}

// deleteCLient - Delete a Client
func (h *handler) deleteCLient(c *gin.Context) {
	id := c.Param(idParam)
	if len(strings.TrimSpace(id)) == 0 {
		c.JSON(http.StatusBadRequest, fmt.Errorf("empty id"))
		return
	}

	err := h.actions.DeleteClient(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusNoContent, nil)
}
