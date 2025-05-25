package handler

import (
	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func (h *handler) healthHandler(c *gin.Context) {
	c.JSON(200, HealthResponse{Status: "ok"})
}
