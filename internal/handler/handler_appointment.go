package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/internal/pkg/dto"
)

func servicesDomainToServicesDTO(servicesDomain []domain.Service) []dto.Service {
	var services = make([]dto.Service, 0)
	for _, s := range servicesDomain {
		services = append(services, dto.Service{
			ID:            s.ID,
			Name:          s.Name,
			Price:         s.Price,
			PaymentType:   s.PaymentType,
			Description:   s.Description,
			ServiceDate:   s.ServiceDate,
			ClientID:      s.Client.ID,
			ClientName:    s.Client.Name,
			AttendantID:   s.Attendant.ID,
			AttendantName: s.Attendant.Name,
		})
	}

	return services
}

func servicesDomainToServiceListResponseDTO(services []domain.Service, csTcsDomain []domain.ClientServiceTypeCount) dto.ServiceListResponse {
	var serviceSummaries = make([]dto.ServiceSummary, 0)
	for _, csc := range csTcsDomain {
		serviceSummaries = append(serviceSummaries, dto.ServiceSummary{
			Name:  csc.ServiceType.Name,
			Count: csc.Count,
		})
	}

	return dto.ServiceListResponse{
		Services:         servicesDomainToServicesDTO(services),
		ServiceSummaries: serviceSummaries,
	}
}

// listServices - List Services
func (h *handler) listServices(c *gin.Context) {
	qps := []domain.Param{}
	for key, value := range c.Request.URL.Query() {
		for _, v := range value {
			qps = append(qps, domain.Param{Key: key, Value: v})
		}
	}

	services, err := h.actions.ListServices(qps)
	if err != nil {
		httpError := newHandlerEror(err)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	if len(services) == 0 {
		c.JSON(http.StatusOK, dto.ServiceListResponse{})
	}

	clientID := services[0].Client.ID
	countOfServiceTypes, err := h.actions.GetClientServicesCount(clientID)
	if err != nil {
		httpError := newHandlerEror(err)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	c.JSON(http.StatusOK, servicesDomainToServiceListResponseDTO(services, countOfServiceTypes))
}
