package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/internal/pkg/dto"
)

func serviceAPIToDomain(s dto.Service) domain.Service {
	return domain.Service{
		ID:          s.ID,
		Client:      clientDTOToDomain(s.Client),
		Attendant:   attendantAPIToDomain(s.Attendant),
		Price:       s.Price,
		ServiceType: s.ServiceType,
		PaymentType: s.PaymentType,
		Description: s.Description,
		ServiceDate: s.ServiceDate,
	}
}

func servicecDomainToAPI(s domain.Service) dto.Service {
	return dto.Service{
		ID:          s.ID,
		Client:      clientDomainToDTO(s.Client),
		Attendant:   attendantDomainToAPI(s.Attendant),
		Price:       s.Price,
		ServiceType: s.ServiceType,
		PaymentType: s.PaymentType,
		Description: s.Description,
		ServiceDate: s.ServiceDate,
	}
}

func serviceTypesCountDomainToAPI(csc domain.ClientServiceTypeCount) dto.ServiceTypeCount {
	return dto.ServiceTypeCount{
		ServiceType: csc.ServiceType.Description,
		Count:       csc.Count,
	}
}

func servicecDomainToAPIList(sList []domain.Service, cscList []domain.ClientServiceTypeCount) dto.ServiceList {
	var serviceList = make([]dto.Service, 0)
	for _, s := range sList {
		serviceList = append(serviceList, servicecDomainToAPI(s))
	}

	var serviceTypesCount = make([]dto.ServiceTypeCount, 0)
	for _, csc := range cscList {
		serviceTypesCount = append(serviceTypesCount, serviceTypesCountDomainToAPI(csc))
	}

	return dto.ServiceList{
		Items:             serviceList,
		ServiceTypesCount: serviceTypesCount,
	}
}

// createService - Create a Service
func (h *handler) createService(c *gin.Context) {
	var serviceAPI dto.Service

	err := c.BindJSON(&serviceAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	service, err := h.actions.CreateService(serviceAPIToDomain(serviceAPI))
	if err != nil {
		httpError := newHandlerEror(err)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	c.JSON(http.StatusCreated, servicecDomainToAPI(service))
}

// listClientServices - List the client Services
func (h *handler) listClientServices(c *gin.Context) {
	clientID := c.Param(idParam)
	if len(strings.TrimSpace(clientID)) == 0 {
		c.JSON(http.StatusBadRequest, fmt.Errorf("empty client id"))
		return
	}

	qps := []domain.Param{}
	for key, value := range c.Request.URL.Query() {
		for _, v := range value {
			qps = append(qps, domain.Param{Key: key, Value: v})
		}
	}

	services, err := h.actions.ListServicesByClient(clientID, qps)
	if err != nil {
		httpError := newHandlerEror(err)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	countOfServiceTypes, err := h.actions.GetClientServicesCount(clientID)
	if err != nil {
		httpError := newHandlerEror(err)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	c.JSON(http.StatusOK, servicecDomainToAPIList(services, countOfServiceTypes))
}
