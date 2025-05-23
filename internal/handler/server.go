// (C) Copyright 2023-2024 Hewlett Packard Enterprise Development LP

package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// ServerInterface is the interface to be implemented by the API handler
type ServerInterface interface {
	//example createResourceProvider(c *gin.Context)

	Routes() Routes
}

var (
	server     *http.Server
	serverOnce sync.Once
)

// ProvideServer creates the HTTP server instance
func ProvideServer(
	apiHandler ServerInterface,
	logger *slog.Logger,
	address string,
	port int,
) *http.Server {
	serverOnce.Do(func() {
		server = newServer(apiHandler, logger, address, port)
	})

	return server
}

func newServer(
	apiHandler ServerInterface,
	logger *slog.Logger,
	address string,
	port int,
) *http.Server {
	logger.Debug("NewServiceRegistryServer")

	var router = NewRouter(
		apiHandler.Routes(),
	)

	// EnableJsonDecoderDisallowUnknownFields enables to reject requests
	// with unknown attributes in the JSON body.
	gin.EnableJsonDecoderDisallowUnknownFields()

	return &http.Server{
		Handler:           router,
		Addr:              fmt.Sprintf("%s:%d", address, port),
		ReadHeaderTimeout: 5 * time.Second,
	}
}
