package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/pkg/apimodel"
)

func serviceAPIToDomain(s apimodel.Service) domain.Service {
	return domain.Service{
		ID:          s.ID,
		Client:      userAPIToDomain(s.Client),
		Attendant:   userAPIToDomain(s.Attendant),
		Price:       s.Price,
		ServiceType: s.ServiceType,
		PaymentType: s.PaymentType,
		Description: s.Description,
	}
}

func servicecDomainToAPI(s domain.Service) apimodel.Service {
	return apimodel.Service{
		ID:          s.ID,
		Client:      userDomainToAPI(s.Client),
		Attendant:   userDomainToAPI(s.Attendant),
		Price:       s.Price,
		ServiceType: s.ServiceType,
		PaymentType: s.PaymentType,
		Description: s.Description,
	}
}

// createService - Create a Service
func (h *handler) createService(c *gin.Context) {
	var serviceAPI apimodel.Service

	err := c.BindJSON(&serviceAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	service, err := h.actions.CreateService(serviceAPIToDomain(serviceAPI))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, servicecDomainToAPI(service))
}
