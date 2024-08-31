package repository

import (
	"github.com/google/uuid"
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/internal/repository/model"
	"gorm.io/gorm"
)

func (db *GormRepository) CreateService(service domain.Service, attendantUUID, clientUUID string) (domain.Service, error) {
	var serviceModel model.Service

	err := db.Transaction(func(tx *gorm.DB) (err error) {
		serviceType, err := db.getServiceType(tx, service.ServiceType)
		if err != nil {
			return err
		}

		serviceModel = model.Service{
			UUID:          uuid.NewString(),
			ClientUUID:    clientUUID,
			AttedantUUID:  attendantUUID,
			Price:         service.Price,
			ServiceTypeID: serviceType.ID,
			PaymentType:   domain.ToPaymentType(service.PaymentType),
			Description:   service.Description,
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
	clientServiceCountModel := model.ClientServiceCount{
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

func (db *GormRepository) updateClientServiceCount(tx *gorm.DB, clientServiceCount model.ClientServiceCount) error {
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

func (db *GormRepository) getClientServiceCount(tx *gorm.DB, cliendUUID string, serviceTypeID uint) (model.ClientServiceCount, error) {
	clientServiceCount := model.ClientServiceCount{}
	err := tx.Where("client_uuid = ? AND service_type_id = ?", cliendUUID, serviceTypeID).Find(&clientServiceCount).Error
	if err != nil {
		return model.ClientServiceCount{}, err
	}
	if clientServiceCount.ClientUUID == "" {
		return model.ClientServiceCount{}, domain.ErrNotFound
	}

	return clientServiceCount, nil
}
