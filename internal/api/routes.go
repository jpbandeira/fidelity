// (C) Copyright 2023-2024 Hewlett Packard Enterprise Development LP

package api

import (
	"net/http"
)

const (
	baseEP string = "/fidelity"

	clientsEP string = baseEP + "/clients"
)

func (h *handler) featureEndpointExampleRoutes() Routes {
	return Routes{
		{
			Name:        "CreateClient",
			Method:      http.MethodPost,
			Pattern:     clientsEP,
			HandlerFunc: h.createClient,
		},
	}
}

// Routes builds the server endpoint route to handlers
func (h *handler) Routes() Routes {
	rts := Routes{}
	rts = append(rts, h.featureEndpointExampleRoutes()...)

	return rts
}
