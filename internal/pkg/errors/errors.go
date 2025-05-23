package domain

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrInvalidParameter = errors.New("invalid parameter")
)

const (
	ErrFormatString string = "%w: %s"
)

type ValidationError struct {
	Field string
	Msg   string
}

func (vErr ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", vErr.Msg, vErr.Field)
}
