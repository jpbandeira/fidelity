// (C) Copyright 2023-2024 Hewlett Packard Enterprise Development LP

package api

import (
	"net/http"
)

const (
	idParam string = "id"

	baseEP string = "/fidelity"

	userEP       string = baseEP + "/users"
	singleUserEP string = userEP + "/:" + idParam

	serviceEP string = baseEP + "/services"
)

func (h *handler) userRoutes() Routes {
	return Routes{
		{
			Name:        "CreateUser",
			Method:      http.MethodPost,
			Pattern:     userEP,
			HandlerFunc: h.createUser,
		},
		{
			Name:        "UpdateUser",
			Method:      http.MethodPut,
			Pattern:     userEP,
			HandlerFunc: h.updateUser,
		},
		{
			Name:        "ListUser",
			Method:      http.MethodGet,
			Pattern:     userEP,
			HandlerFunc: h.listUser,
		},
		{
			Name:        "DeleteUser",
			Method:      http.MethodDelete,
			Pattern:     singleUserEP,
			HandlerFunc: h.deleteUser,
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
	}
}

// Routes builds the server endpoint route to handlers
func (h *handler) Routes() Routes {
	rts := Routes{}
	rts = append(rts, h.userRoutes()...)
	rts = append(rts, h.serviceRoutes()...)

	return rts
}
