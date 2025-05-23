package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/internal/pkg/apimodel"
)

func serviceTypeAPIToDomain(s apimodel.ServiceType) domain.ServiceType {
	return domain.ServiceType{
		Description: s.Description,
	}
}

func serviceTypeDomainToAPI(s domain.ServiceType) apimodel.ServiceType {
	return apimodel.ServiceType{
		Description: s.Description,
	}
}

// createServiceType - Create a ServiceType
// func (h *handler) createServiceType(c *gin.Context) {
// 	var serviceTypeAPI apimodel.ServiceType

// 	err := c.BindJSON(&serviceTypeAPI)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	serviceType, err := h.actions.CreateServiceType(serviceTypeTypeAPIToDomain(serviceTypeAPI))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusCreated, serviceTypecDomainToAPI(serviceType))
// }

// listServiceType - List a list of ServiceTypes
func (h *handler) listServiceType(c *gin.Context) {
	qps := []domain.Param{}
	for key, value := range c.Request.URL.Query() {
		for _, v := range value {
			qps = append(qps, domain.Param{Key: key, Value: v})
		}
	}

	serviceTypes, err := h.actions.ListServiceTypes(qps)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	var result = make([]apimodel.ServiceType, 0)
	for _, st := range serviceTypes {
		result = append(result, serviceTypeDomainToAPI(st))
	}

	c.JSON(http.StatusOK, result)
}
