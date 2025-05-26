package repository

import (
	"github.com/google/uuid"
	"github.com/jp/fidelity/internal/domain"
	ferros "github.com/jp/fidelity/internal/pkg/errors"
	"github.com/jp/fidelity/internal/repository/model"
	"gorm.io/gorm"
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

func (db *GormRepository) CreateAppointment(appt domain.Appointment) (domain.Appointment, error) {
	apptUUID := uuid.NewString()
	modelAppointment := model.Appointment{
		UUID:          apptUUID,
		ClientUUID:    appt.Client.ID,
		AttendantUUID: appt.Attendant.ID,
		Client:        model.ClientDomainToRepo(appt.Client),
		Attendant:     model.AttendantDomainToRepo(appt.Attendant),
	}

	services := make([]model.Service, 0, len(appt.Services))
	for _, s := range appt.Services {
		stype, err := db.getServiceType(db.DB, s.Name)
		if err != nil {
			return domain.Appointment{}, err
		}

		services = append(services, model.Service{
			UUID:            uuid.NewString(),
			AppointmentUUID: apptUUID,
			ServiceTypeID:   stype.ID,
			ServiceType:     stype,
			Price:           s.Price,
			PaymentType:     domain.ToPaymentType(s.PaymentType),
			Description:     s.Description,
			ServiceDate:     s.ServiceDate,
		})
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		modelAppointment.Services = services
		if err := tx.Create(&modelAppointment).Error; err != nil {
			return err
		}

		for _, s := range modelAppointment.Services {
			csc, err := db.getClientServiceCount(tx, appt.Client.ID, s.ServiceTypeID)
			if err != nil && err != ferros.ErrNotFound {
				return err
			}

			if csc.ClientUUID != "" {
				if err := db.updateClientServiceCount(tx, csc); err != nil {
					return err
				}
			} else {
				if err := db.createClientServiceCount(tx, s.ServiceTypeID, appt.Client.ID); err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		return domain.Appointment{}, err
	}

	return modelAppointment.RepoToDomain(), nil
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

func (db *GormRepository) getClientServiceCount(tx *gorm.DB, cliendUUID string, serviceTypeID uint) (model.ClientServiceTypeCount, error) {
	clientServiceCount := model.ClientServiceTypeCount{}
	err := tx.Where("client_uuid = ? AND service_type_id = ?", cliendUUID, serviceTypeID).Find(&clientServiceCount).Error
	if err != nil {
		return model.ClientServiceTypeCount{}, err
	}
	if clientServiceCount.ClientUUID == "" {
		return model.ClientServiceTypeCount{}, ferros.ErrNotFound
	}

	return clientServiceCount, nil
}

func (db *GormRepository) updateClientServiceCount(tx *gorm.DB, clientServiceCount model.ClientServiceTypeCount) error {
	clientServiceCount.ServiceCount = clientServiceCount.ServiceCount + 1
	err := tx.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Select("*").
		Where("client_uuid = ? AND service_type_id = ?", clientServiceCount.ClientUUID, clientServiceCount.ServiceTypeID).
		Updates(&clientServiceCount).Error
	if err != nil {
		return err
	}

	return nil
}

func (db *GormRepository) createClientServiceCount(tx *gorm.DB, serviceTypeID uint, clientUUID string) error {
	clientServiceCountModel := model.ClientServiceTypeCount{
		ServiceTypeID: serviceTypeID,
		ClientUUID:    clientUUID,
		ServiceCount:  1,
	}

	err := tx.Create(&clientServiceCountModel).Error
	if err != nil {
		return err
	}

	return nil
}

func (db *GormRepository) getServiceType(tx *gorm.DB, serviceType string) (model.ServiceType, error) {
	var serviceTypeModel = model.ServiceType{}
	t := tx.Where("name = ?", serviceType).Find(&serviceTypeModel)
	if t.Error != nil {
		return model.ServiceType{}, t.Error
	}
	if serviceTypeModel.ID == 0 {
		return model.ServiceType{}, ferros.ErrNotFound
	}

	return serviceTypeModel, nil
}
