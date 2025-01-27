// (C) Copyright 2023-2024 Hewlett Packard Enterprise Development LP

package api

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

	serviceEP        string = baseEP + "/services"
	clientServicesRP string = singleClientEP + "/services"
	serviceTypeEP    string = baseEP + "/service-types"
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

func (h *handler) attendantRoutes() Routes {
	return Routes{
		{
			Name:        "CreateAttendant",
			Method:      http.MethodPost,
			Pattern:     attendantEP,
			HandlerFunc: h.createAttendant,
		},
		{
			Name:        "UpdateAttendant",
			Method:      http.MethodPut,
			Pattern:     attendantEP,
			HandlerFunc: h.updateAttendant,
		},
		{
			Name:        "ListAttendant",
			Method:      http.MethodGet,
			Pattern:     attendantEP,
			HandlerFunc: h.listAttendant,
		},
		{
			Name:        "DeleteAttendant",
			Method:      http.MethodDelete,
			Pattern:     singleAttendantEP,
			HandlerFunc: h.deleteAttendant,
		},
	}
}

func (h *handler) serviceRoutes() Routes {
	return Routes{
		{
			Name:        "CreateService",
			Method:      http.MethodPost,
			Pattern:     serviceEP,
			HandlerFunc: h.createService,
		},
		{
			Name:        "ListClientServices",
			Method:      http.MethodGet,
			Pattern:     clientServicesRP,
			HandlerFunc: h.listClientServices,
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
	}
}

// Routes builds the server endpoint route to handlers
func (h *handler) Routes() Routes {
	rts := Routes{}
	rts = append(rts, h.clientRoutes()...)
	rts = append(rts, h.attendantRoutes()...)
	rts = append(rts, h.serviceRoutes()...)
	rts = append(rts, h.serviceTypeRoutes()...)

	return rts
}
