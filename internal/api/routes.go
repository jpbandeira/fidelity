// (C) Copyright 2023-2024 Hewlett Packard Enterprise Development LP

package api

import (
	"net/http"
)

const (
	baseEP string = "/fidelity"

	personsEP string = baseEP + "/persons"
)

func (h *handler) featureEndpointExampleRoutes() Routes {
	return Routes{
		{
			Name:        "CreatePerson",
			Method:      http.MethodPost,
			Pattern:     personsEP,
			HandlerFunc: h.createPerson,
		},
	}
}

// Routes builds the server endpoint route to handlers
func (h *handler) Routes() Routes {
	rts := Routes{}
	rts = append(rts, h.featureEndpointExampleRoutes()...)

	return rts
}
