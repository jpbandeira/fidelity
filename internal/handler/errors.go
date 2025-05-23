package handler

import (
	"errors"
	"net/http"

	ferros "github.com/jp/fidelity/internal/pkg/errors"
)

type HandlerError struct {
	msg         string
	error_type  string
	status_code int
}

func newHandlerEror(err error) HandlerError {
	var errorType *ferros.ValidationError
	if errors.As(err, &errorType) {
		if errors.Is(err, ferros.ErrInvalidParameter) {
			return HandlerError{
				msg:         errorType.Error(),
				error_type:  ferros.ErrInvalidParameter.Error(),
				status_code: http.StatusBadRequest,
			}
		}

		if errors.Is(err, ferros.ErrNotFound) {
			return HandlerError{
				msg:         errorType.Error(),
				error_type:  ferros.ErrNotFound.Error(),
				status_code: http.StatusNotFound,
			}
		}
	}

	return HandlerError{
		msg:         err.Error(),
		error_type:  "Internal Server error",
		status_code: http.StatusInternalServerError,
	}
}
