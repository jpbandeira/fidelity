package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/internal/pkg/dto"
	ferros "github.com/jp/fidelity/internal/pkg/errors"
)

func attendantAPIToDomain(c dto.Attendant) domain.Attendant {
	return domain.Attendant{
		ID:    c.ID,
		Name:  c.Name,
		Email: c.Email,
		Phone: c.Phone,
	}
}

func attendantDomainToAPI(c domain.Attendant) dto.Attendant {
	return dto.Attendant{
		ID:    c.ID,
		Name:  c.Name,
		Email: c.Email,
		Phone: c.Phone,
	}
}

// createAttendant - Create a Attendant
func (h *handler) createAttendant(c *gin.Context) {
	var attendantAPI dto.Attendant

	err := c.BindJSON(&attendantAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	attendant, err := h.actions.CreateAttendant(attendantAPIToDomain(attendantAPI))
	if err != nil {
		httpError := newHandlerEror(err)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	c.JSON(http.StatusCreated, attendantDomainToAPI(attendant))
}

// updateAttendant - Update a Attendant
func (h *handler) updateAttendant(c *gin.Context) {
	var attendantAPI dto.Attendant

	err := c.BindJSON(&attendantAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	attendant, err := h.actions.UpdateAttendant(attendantAPIToDomain(attendantAPI))
	if err != nil {
		httpError := newHandlerEror(err)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	c.JSON(http.StatusOK, attendantDomainToAPI(attendant))
}

// listAttendant - List a list of Attendant
func (h *handler) listAttendant(c *gin.Context) {
	qps := []domain.Param{}
	for key, value := range c.Request.URL.Query() {
		for _, v := range value {
			qps = append(qps, domain.Param{Key: key, Value: v})
		}
	}

	attendants, err := h.actions.ListAttendants(qps)
	if err != nil {
		httpError := newHandlerEror(err)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	var result = make([]dto.Attendant, 0)
	for _, att := range attendants {
		result = append(result, attendantDomainToAPI(att))
	}

	c.JSON(http.StatusOK, result)
}

// deleteAttendant - Delete a Attendant
func (h *handler) deleteAttendant(c *gin.Context) {
	id := c.Param(idParam)
	if len(strings.TrimSpace(id)) == 0 {
		c.JSON(http.StatusBadRequest, fmt.Errorf("%w: %s", ferros.ErrInvalidParameter, "empty id"))
		return
	}

	err := h.actions.DeleteAttendant(id)
	if err != nil {
		httpError := newHandlerEror(err)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
