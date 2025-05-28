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
		Client:        model.ClientDomainToRepo(appt.Client),
		AttendantUUID: appt.AttendantID,
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
			serviceSummary, err := db.getServiceSummary(tx, appt.Client.ID, s.ServiceTypeID)
			if err != nil && err != ferros.ErrNotFound {
				return err
			}

			if serviceSummary.ClientUUID != "" {
				serviceSummary.Count = serviceSummary.Count + 1
				serviceSummary.TotalPrice = serviceSummary.TotalPrice + s.Price

				if err := db.updateServiceSummary(tx, serviceSummary); err != nil {
					return err
				}
			} else {
				serviceSummary := model.ServiceSummary{
					ServiceTypeID: s.ServiceTypeID,
					ClientUUID:    modelAppointment.ClientUUID,
					Count:         1,
					TotalPrice:    s.Price,
				}
				err := db.createServiceSummary(tx, serviceSummary)
				if err != nil {
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

func (db *GormRepository) GetClientServicesCount(cliendUUID string) ([]domain.ServiceSummary, error) {
	clientServicesCount := []model.ServiceSummary{}
	err := db.Preload("Client").Preload("ServiceType").Where("client_uuid = ?", cliendUUID).Find(&clientServicesCount).Error
	if err != nil {
		return []domain.ServiceSummary{}, err
	}

	var result []domain.ServiceSummary
	for _, csc := range clientServicesCount {
		result = append(result, csc.RepoToDomain())
	}

	return result, nil
}

func (db *GormRepository) getServiceSummary(tx *gorm.DB, cliendUUID string, serviceTypeID uint) (model.ServiceSummary, error) {
	serviceSummary := model.ServiceSummary{}
	err := tx.Where("client_uuid = ? AND service_type_id = ?", cliendUUID, serviceTypeID).Find(&serviceSummary).Error
	if err != nil {
		return model.ServiceSummary{}, err
	}
	if serviceSummary.ClientUUID == "" {
		return model.ServiceSummary{}, ferros.ErrNotFound
	}

	return serviceSummary, nil
}

func (db *GormRepository) updateServiceSummary(tx *gorm.DB, serviceSummary model.ServiceSummary) error {
	err := tx.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Select("*").
		Where("client_uuid = ? AND service_type_id = ?", serviceSummary.ClientUUID, serviceSummary.ServiceTypeID).
		Updates(&serviceSummary).Error
	if err != nil {
		return err
	}

	return nil
}

func (db *GormRepository) createServiceSummary(tx *gorm.DB, serviceSummary model.ServiceSummary) error {
	err := tx.Create(&serviceSummary).Error
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
