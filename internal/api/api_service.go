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
		Client:      clientAPIToDomain(s.Client),
		Attendant:   attendantAPIToDomain(s.Attendant),
		Price:       s.Price,
		ServiceType: s.ServiceType,
		PaymentType: s.PaymentType,
		Description: s.Description,
		ServiceDate: s.ServiceDate,
	}
}

func servicecDomainToAPI(s domain.Service) apimodel.Service {
	return apimodel.Service{
		ID:          s.ID,
		Client:      clientDomainToAPI(s.Client),
		Attendant:   attendantDomainToAPI(s.Attendant),
		Price:       s.Price,
		ServiceType: s.ServiceType,
		PaymentType: s.PaymentType,
		Description: s.Description,
		ServiceDate: s.ServiceDate,
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

// listService - List a list of Services
func (h *handler) listService(c *gin.Context) {
	qps := []domain.Param{}
	for key, value := range c.Request.URL.Query() {
		for _, v := range value {
			qps = append(qps, domain.Param{Key: key, Value: v})
		}
	}

	users, err := h.actions.ListServices(qps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, users)
}
