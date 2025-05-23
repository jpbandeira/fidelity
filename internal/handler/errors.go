package handler

import (
	"errors"
	"net/http"

	ferros "github.com/jp/fidelity/internal/pkg/errors"
)

type HandlerError struct {
	Msg        string `json:"msg"`
	ErrorType  string `json:"errorType"`
	StatusCode int    `json:"-"`
}

func newHandlerEror(err error) HandlerError {
	var validationErr *ferros.ValidationError
	if errors.As(err, &validationErr) {
		if errors.Is(err, ferros.ErrInvalidParameter) {
			return HandlerError{
				Msg:        validationErr.Error(),
				ErrorType:  ferros.ErrInvalidParameter.Error(),
				StatusCode: http.StatusBadRequest,
			}
		}
	}

	var nfErr *ferros.NotFoundError
	if errors.As(err, &nfErr) {
		return HandlerError{
			Msg:        nfErr.Error(),
			ErrorType:  ferros.ErrNotFound.Error(),
			StatusCode: http.StatusNotFound,
		}
	}

	return HandlerError{
		Msg:        err.Error(),
		ErrorType:  ferros.ErrInternalServerError.Error(),
		StatusCode: http.StatusInternalServerError,
	}
}
