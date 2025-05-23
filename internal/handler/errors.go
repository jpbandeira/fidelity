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
	var validationErr *ferros.ValidationError
	if errors.As(err, &validationErr) {
		if errors.Is(err, ferros.ErrInvalidParameter) {
			return HandlerError{
				msg:         validationErr.Error(),
				error_type:  ferros.ErrInvalidParameter.Error(),
				status_code: http.StatusBadRequest,
			}
		}
	}

	var nfErr *ferros.NotFoundError
	if errors.As(err, &nfErr) {
		return HandlerError{
			msg:         nfErr.Error(),
			error_type:  ferros.ErrNotFound.Error(),
			status_code: http.StatusNotFound,
		}
	}

	return HandlerError{
		msg:         err.Error(),
		error_type:  ferros.ErrInternalServerError.Error(),
		status_code: http.StatusInternalServerError,
	}
}
