package handler

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jp/fidelity/internal/domain"
// 	"github.com/jp/fidelity/internal/pkg/dto"
// )

// func serviceBatchDTOToDomain(s dto.ServiceBatch) domain.ServiceBatch {
// 	var services = make([]domain.Service, 0)
// 	for _, v := range s.Items {
// 		services = append(services, domain.Service{
// 			ID:          v.ID,
// 			Price:       v.Price,
// 			ServiceType: v.ServiceType,
// 			PaymentType: v.PaymentType,
// 			Description: v.Description,
// 			ServiceDate: v.ServiceDate,
// 		})
// 	}

// 	return domain.ServiceBatch{
// 		Client:    clientDTOToDomain(s.Client),
// 		Attendant: attendantDTOToDomain(s.Attendant),
// 		Items:     services,
// 	}
// }

// func serviceDomainToDTO(s domain.Service) dto.Service {
// 	return dto.Service{
// 		ID:          s.ID,
// 		Price:       s.Price,
// 		ServiceType: s.ServiceType,
// 		PaymentType: s.PaymentType,
// 		Description: s.Description,
// 		ServiceDate: s.ServiceDate,
// 	}
// }

// func serviceTypesCountDomainToDTO(csc domain.ClientServiceTypeCount) dto.ServiceTypeCount {
// 	return dto.ServiceTypeCount{
// 		ServiceType: csc.ServiceType.Description,
// 		Count:       csc.Count,
// 	}
// }

// func serviceBatchDomainToDTO(sBatch domain.ServiceBatch) dto.ServiceBatch {
// 	var services = make([]dto.Service, 0)
// 	for _, s := range sBatch.Items {
// 		services = append(services, serviceDomainToDTO(s))
// 	}

// 	return dto.ServiceBatch{
// 		Client:    clientDomainToDTO(sBatch.Client),
// 		Attendant: attendantDomainToDTO(sBatch.Attendant),
// 		Items:     services,
// 	}
// }

// func serviceBatchToServiceListDTO(sBatch dto.ServiceBatch, csTcList []domain.ClientServiceTypeCount) dto.ServiceList {
// 	var serviceTypesCount = make([]dto.ServiceTypeCount, 0)
// 	for _, csc := range csTcList {
// 		serviceTypesCount = append(serviceTypesCount, serviceTypesCountDomainToDTO(csc))
// 	}

// 	return dto.ServiceList{
// 		Item:              sBatch,
// 		ServiceTypesCount: serviceTypesCount,
// 	}
// }



// // listServices - List Services
// func (h *handler) listServices(c *gin.Context) {
// 	qps := []domain.Param{}
// 	for key, value := range c.Request.URL.Query() {
// 		for _, v := range value {
// 			qps = append(qps, domain.Param{Key: key, Value: v})
// 		}
// 	}

// 	serviceBatchDomain, err := h.actions.ListServices(qps)
// 	if err != nil {
// 		httpError := newHandlerEror(err)
// 		c.JSON(httpError.StatusCode, httpError)
// 		return
// 	}

// 	serviceBatch := serviceBatchDomainToDTO(serviceBatchDomain)

// 	countOfServiceTypes, err := h.actions.GetClientServicesCount(serviceBatch.Client.ID)
// 	if err != nil {
// 		httpError := newHandlerEror(err)
// 		c.JSON(httpError.StatusCode, httpError)
// 		return
// 	}

// 	c.JSON(http.StatusOK, serviceBatchToServiceListDTO(serviceBatch, countOfServiceTypes))
// }
