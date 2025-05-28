// (C) Copyright 2023-2024 Hewlett Packard Enterprise Development LP

package handler

import (
	"net/http"
)

const (
	idParam string = "id"

	baseEP string = "/fidelity"

	clientEP          string = baseEP + "/clients"
	attendantEP       string = baseEP + "/attendants"
	singleClientEP    string = clientEP + "/:" + idParam
	singleAttendantEP string = attendantEP + "/:" + idParam

	appointmentEP string = baseEP + "/appointments"
	serviceEP     string = appointmentEP + "/services"
	serviceTypeEP string = baseEP + "/service-types"
)

func (h *handler) clientRoutes() Routes {
	return Routes{
		{
			Name:        "CreateClient",
			Method:      http.MethodPost,
			Pattern:     clientEP,
			HandlerFunc: h.createClient,
		},
		{
			Name:        "UpdateClient",
			Method:      http.MethodPut,
			Pattern:     clientEP,
			HandlerFunc: h.updateClient,
		},
		{
			Name:        "ListClient",
			Method:      http.MethodGet,
			Pattern:     clientEP,
			HandlerFunc: h.listClient,
		},
		{
			Name:        "DeleteClient",
			Method:      http.MethodDelete,
			Pattern:     singleClientEP,
			HandlerFunc: h.deleteCLient,
		},
	}
}

func (h *handler) serviceRoutes() Routes {
	return Routes{
		{
			Name:        "CreateAppointment",
			Method:      http.MethodPost,
			Pattern:     appointmentEP,
			HandlerFunc: h.createAppointment,
		},
		{
			Name:        "ListServices",
			Method:      http.MethodGet,
			Pattern:     serviceEP,
			HandlerFunc: h.listServices,
		},
	}
}

func (h *handler) serviceTypeRoutes() Routes {
	return Routes{
		{
			Name:        "ListServiceType",
			Method:      http.MethodGet,
			Pattern:     serviceTypeEP,
			HandlerFunc: h.listServiceType,
		},
		{
			Name:        "CreateServiceType",
			Method:      http.MethodPost,
			Pattern:     serviceTypeEP,
			HandlerFunc: h.createServiceType,
		},
	}
}

func (h *handler) healthRoutes() Routes {
	return Routes{
		{
			Name:        "Health",
			Method:      http.MethodGet,
			Pattern:     "/health",
			HandlerFunc: h.healthHandler,
		},
	}
}

// Routes builds the server endpoint route to handlers
func (h *handler) Routes() Routes {
	rts := Routes{}
	rts = append(rts, h.clientRoutes()...)
	rts = append(rts, h.serviceRoutes()...)
	rts = append(rts, h.serviceTypeRoutes()...)
	rts = append(rts, h.healthRoutes()...)

	return rts
}
