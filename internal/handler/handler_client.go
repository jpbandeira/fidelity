package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/internal/pkg/dto"
)

func clientDTOToDomain(c dto.Client) domain.Client {
	return domain.Client{
		ID:    c.ID,
		Name:  c.Name,
		Email: c.Email,
		Phone: c.Phone,
	}
}

func clientDomainToDTO(c domain.Client) dto.Client {
	return dto.Client{
		ID:    c.ID,
		Name:  c.Name,
		Email: c.Email,
		Phone: c.Phone,
	}
}

// createClient - Create a Client
func (h *handler) createClient(c *gin.Context) {
	var clientAPI dto.Client

	err := c.BindJSON(&clientAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client, err := h.actions.CreateClient(clientDTOToDomain(clientAPI))
	if err != nil {
		httpError := newHandlerEror(err)
		c.JSON(httpError.status_code, httpError)
		return
	}

	c.JSON(http.StatusCreated, clientDomainToDTO(client))
}

// updateClient - Update a Client
func (h *handler) updateClient(c *gin.Context) {
	var clientAPI dto.Client

	err := c.BindJSON(&clientAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client, err := h.actions.UpdateClient(clientDTOToDomain(clientAPI))
	if err != nil {
		httpError := newHandlerEror(err)
		c.JSON(httpError.status_code, httpError)
		return
	}

	c.JSON(http.StatusOK, clientDomainToDTO(client))
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
		httpError := newHandlerEror(err)
		c.JSON(httpError.status_code, httpError)
		return
	}

	var result = make([]dto.Client, 0)
	for _, c := range clients {
		result = append(result, clientDomainToDTO(c))
	}

	c.JSON(http.StatusOK, result)
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
		httpError := newHandlerEror(err)
		c.JSON(httpError.status_code, httpError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
