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

func dtoToDomainAppointment(appt dto.Appointment) domain.Appointment {
	services := make([]domain.Service, 0, len(appt.Services))
	for _, s := range appt.Services {
		services = append(services, domain.Service{
			ID:          s.ID,
			Name:        s.Name,
			Price:       s.Price,
			PaymentType: s.PaymentType,
			Description: s.Description,
			ServiceDate: s.ServiceDate,
			Client: domain.Client{
				ID:   s.ClientID,
				Name: s.ClientName,
			},
			Attendant: domain.Attendant{
				ID:   s.AttendantID,
				Name: s.AttendantName,
			},
		})
	}

	return domain.Appointment{
		ID: appt.ID,
		Client: domain.Client{
			ID:   appt.Client.ID,
			Name: appt.Client.Name,
		},
		Attendant: domain.Attendant{
			ID:   appt.Attendant.ID,
			Name: appt.Attendant.Name,
		},
		Services: services,
	}
}

func domainToDTOAppointment(appt domain.Appointment) dto.Appointment {
	services := make([]dto.Service, 0, len(appt.Services))
	for _, s := range appt.Services {
		services = append(services, dto.Service{
			ID:          s.ID,
			Name:        s.Name,
			Price:       s.Price,
			PaymentType: s.PaymentType,
			Description: s.Description,
			ServiceDate: s.ServiceDate,
		})
	}

	return dto.Appointment{
		ID: appt.ID,
		Client: dto.Client{
			ID:    appt.Client.ID,
			Name:  appt.Client.Name,
			Email: appt.Client.Email,
			Phone: appt.Client.Phone,
		},
		Attendant: dto.Attendant{
			ID:    appt.Attendant.ID,
			Name:  appt.Attendant.Name,
			Email: appt.Attendant.Email,
			Phone: appt.Attendant.Phone,
		},
		Services: services,
	}
}

func (h *handler) createAppointment(c *gin.Context) {
	var apptDTO dto.Appointment
	if err := c.ShouldBindJSON(&apptDTO); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	appt, err := h.actions.CreateAppointment(dtoToDomainAppointment(apptDTO))
	if err != nil {
		httpError := newHandlerEror(err)
		c.JSON(httpError.StatusCode, httpError)
		return
	}

	c.JSON(http.StatusCreated, domainToDTOAppointment(appt))
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
		c.JSON(http.StatusOK, dto.ServiceListResponse{
			Services:         []dto.Service{},
			ServiceSummaries: []dto.ServiceSummary{},
		})
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
