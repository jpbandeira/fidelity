package repository

import (
	"github.com/google/uuid"
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/internal/repository/model"
	"gorm.io/gorm"
)

func (db *GormRepository) CreateService(service domain.Service) (domain.Service, error) {
	var serviceModel model.Service

	err := db.Transaction(func(tx *gorm.DB) (err error) {
		serviceType, err := db.getServiceType(tx, service.ServiceType)
		if err != nil {
			return err
		}

		serviceModel = model.Service{
			UUID:          uuid.NewString(),
			ClientUUID:    service.Client.ID,
			Client:        model.ClientDomainToRepo(service.Client),
			AttedantUUID:  service.Attendant.ID,
			Attendant:     model.AttendantDomainToRepo(service.Attendant),
			Price:         service.Price,
			ServiceTypeID: serviceType.ID,
			ServiceType:   serviceType,
			PaymentType:   domain.ToPaymentType(service.PaymentType),
			Description:   service.Description,
			ServiceDate:   service.ServiceDate,
		}

		err = db.Create(&serviceModel).Error
		if err != nil {
			return err
		}

		clientServiceCount, cscErr := db.getClientServiceCount(tx, serviceModel.ClientUUID, serviceModel.ServiceTypeID)
		if cscErr != nil && cscErr != domain.ErrNotFound {
			return cscErr
		}

		if clientServiceCount.ClientUUID != "" {
			return db.updateClientServiceCount(tx, clientServiceCount)
		}

		return db.createClientServiceCount(tx, serviceModel.ServiceTypeID, serviceModel.ClientUUID)
	})
	if err != nil {
		return domain.Service{}, err
	}

	return serviceModel.RepoToDomain(), nil
}

func (db *GormRepository) getServiceType(tx *gorm.DB, serviceType string) (model.ServiceType, error) {
	var serviceTypeModel = model.ServiceType{}
	t := tx.Where("description = ?", serviceType).Find(&serviceTypeModel)
	if t.Error != nil {
		return model.ServiceType{}, t.Error
	}
	if serviceTypeModel.ID == 0 {
		return model.ServiceType{}, domain.ErrNotFound
	}

	return serviceTypeModel, nil
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

func (db *GormRepository) getClientServiceCount(tx *gorm.DB, cliendUUID string, serviceTypeID uint) (model.ClientServiceTypeCount, error) {
	clientServiceCount := model.ClientServiceTypeCount{}
	err := tx.Where("client_uuid = ? AND service_type_id = ?", cliendUUID, serviceTypeID).Find(&clientServiceCount).Error
	if err != nil {
		return model.ClientServiceTypeCount{}, err
	}
	if clientServiceCount.ClientUUID == "" {
		return model.ClientServiceTypeCount{}, domain.ErrNotFound
	}

	return clientServiceCount, nil
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

func (db *GormRepository) ListServicesByClient(clientID string, params []domain.Param) ([]domain.Service, error) {
	var services []model.Service
	var q string
	var args []interface{}

	for _, v := range params {
		q = q + v.Key + " = ?"
		args = append(args, v.Value)
	}

	q = q + "client_uuid" + " = ?"
	args = append(args, clientID)

	err := db.Preload("Client").Preload("Attendant").Preload("ServiceType").Where(q, args...).Order("service_date DESC").Find(&services).Error
	if err != nil {
		return []domain.Service{}, err
	}

	var result []domain.Service
	for _, s := range services {
		result = append(result, s.RepoToDomain())
	}

	return result, nil
}
