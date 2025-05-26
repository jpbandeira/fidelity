package domain

import (
	"fmt"

	ferros "github.com/jp/fidelity/internal/pkg/errors"
)

type ServiceType struct {
	Name string
}

type ServiceTypeList struct {
	Items []ServiceType
	Total int
	Count int
}

const (
	ServiceTypeEntity string = "sertice type"
)

func (st ServiceType) validateServiceType() error {
	if st.Name == "" {
		return fmt.Errorf(ferros.ErrFormatString, ferros.ErrInvalidParameter, &ferros.ValidationError{
			Field:  ferros.DescriptionField,
			Msg:    ferros.EmptyErrorString,
			Entity: ServiceTypeEntity,
		})
	}

	return nil
}

func (a *actions) CreateServiceType(st ServiceType) (ServiceType, error) {
	err := st.validateServiceType()
	if err != nil {
		return ServiceType{}, err
	}

	return a.db.CreateServiceType(st)
}

func (a actions) ListServiceTypes(params []Param) ([]ServiceType, error) {
	return a.db.ListServiceTypes(params)
}
