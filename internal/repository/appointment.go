package repository

import (
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/internal/repository/model"
)

func (db *GormRepository) ListServices(params []domain.Param) ([]domain.Service, error) {
	var services []model.Service
	var q string
	var args []any

	for i, v := range params {
		if i > 0 {
			q += " AND "
		}
		q += v.Key + " = ?"
		args = append(args, v.Value)
	}

	err := db.
		Joins("JOIN appointments ON appointments.uuid = services.appointment_uuid").
		Preload("Appointment.Client").
		Preload("Appointment.Attendant").
		Preload("ServiceType").
		Where(q, args...).
		Order("service_date DESC").
		Find(&services).Error
	if err != nil {
		return []domain.Service{}, err
	}

	if len(services) == 0 {
		return []domain.Service{}, nil
	}

	return model.ServiceRepoToDomain(services), nil
}

func (db *GormRepository) GetClientServicesCount(cliendUUID string) ([]domain.ClientServiceTypeCount, error) {
	clientServicesCount := []model.ClientServiceTypeCount{}
	err := db.Preload("Client").Preload("ServiceType").Where("client_uuid = ?", cliendUUID).Find(&clientServicesCount).Error
	if err != nil {
		return []domain.ClientServiceTypeCount{}, err
	}

	var result []domain.ClientServiceTypeCount
	for _, csc := range clientServicesCount {
		result = append(result, csc.RepoToDomain())
	}

	return result, nil
}
