package handler

import (
	"log/slog"
	"sync"

	"github.com/jp/fidelity/internal/domain"
)

type handler struct {
	actions domain.Actions
	logger  *slog.Logger
}

var (
	hdl     *handler
	hdlOnce sync.Once
)

// ProvideHandler create the single handler instance that implements the domain actions
func ProvideHandler(actions domain.Actions, logger *slog.Logger) *handler {
	hdlOnce.Do(func() {
		hdl = &handler{
			actions: actions,
			logger:  logger,
		}
	})
	return hdl
}
