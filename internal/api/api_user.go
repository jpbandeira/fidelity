package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/pkg/apimodel"
)

func apiToDomain(c apimodel.User) domain.User {
	return domain.User{
		ID:    c.ID,
		Name:  c.Name,
		Email: c.Email,
		Phone: c.Phone,
	}
}

func domainToAPI(c domain.User) apimodel.User {
	return apimodel.User{
		ID:    c.ID,
		Name:  c.Name,
		Email: c.Email,
		Phone: c.Phone,
	}
}

// createUser - Create a User
func (h *handler) createUser(c *gin.Context) {
	var userAPI apimodel.User

	err := c.BindJSON(&userAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := h.actions.CreateUser(apiToDomain(userAPI))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, domainToAPI(user))
}

// updateUser - Update a User
func (h *handler) updateUser(c *gin.Context) {
	var userAPI apimodel.User

	err := c.BindJSON(&userAPI)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := h.actions.UpdateUser(apiToDomain(userAPI))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, domainToAPI(user))
}

// listUser - List a list of User
func (h *handler) listUser(c *gin.Context) {
	users, err := h.actions.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, users)
}

// deleteUser - Delete a User
func (h *handler) deleteUser(c *gin.Context) {
	id := c.Param(idParam)
	if len(strings.TrimSpace(id)) == 0 {
		c.JSON(http.StatusBadRequest, fmt.Errorf("empty id"))
		return
	}

	err := h.actions.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusNoContent, nil)
}
