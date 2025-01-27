package api

import (
	"fmt"
	"net/http"
	"strings"

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

func clientServiceTypeToAPI(csc domain.ClientServiceCount) apimodel.ClientServiceType {
	return apimodel.ClientServiceType{
		ServiceType: csc.ServiceType.Description,
		Count:       csc.ServiceCount,
	}
}

func servicecDomainToAPIList(sList []domain.Service, cscList []domain.ClientServiceCount) apimodel.ServiceList {
	var apiServiceList = make([]apimodel.Service, 0)
	for _, s := range sList {
		apiServiceList = append(apiServiceList, servicecDomainToAPI(s))
	}

	var clientServiceTypeList = make([]apimodel.ClientServiceType, 0)
	for _, csc := range cscList {
		clientServiceTypeList = append(clientServiceTypeList, clientServiceTypeToAPI(csc))
	}

	return apimodel.ServiceList{
		Items:               apiServiceList,
		CountOfServiceTypes: clientServiceTypeList,
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
		c.JSON(http.StatusInternalServerError, err)
	}

	countOfServiceTypes, err := h.actions.GetClientServicesCount(clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, servicecDomainToAPIList(services, countOfServiceTypes))
}
