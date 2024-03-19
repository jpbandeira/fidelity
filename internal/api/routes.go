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

// Routes builds the server endpoint route to handlers
func (h *handler) Routes() Routes {
	rts := Routes{}
	rts = append(rts, h.userRoutes()...)

	return rts
}
