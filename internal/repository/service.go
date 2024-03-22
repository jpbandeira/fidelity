package repository

import (
	"github.com/google/uuid"
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/internal/repository/model"
	"gorm.io/gorm"
)

func (db *GormRepository) CreateService(service domain.Service) (domain.Service, error) {
	var serviceModel model.Service

	err := db.Transaction(func(tx *gorm.DB) error {
		client, err := db.getUser(tx, service.Client.ID)
		if err != nil {
			return err
		}

		attendant, err := db.getUser(tx, service.Attendant.ID)
		if err != nil {
			return err
		}

		var serviceType model.ServiceType
		t := db.Where("description = ?", service.ServiceType).Find(&serviceType)
		if t.Error != nil {
			return err
		}
		if serviceType.ID == 0 {
			return domain.ErrNotFound
		}

		serviceModel = model.Service{
			UUID:          uuid.NewString(),
			ClientID:      client.ID,
			AttedantID:    attendant.ID,
			Price:         service.Price,
			ServiceTypeID: serviceType.ID,
			PaymentType:   domain.ToPaymentType(service.PaymentType),
			Description:   service.Description,
		}

		err = db.Create(&serviceModel).Error
		if err != nil {
			return err
		}

		clientServiceCount, err := db.getClientServiceCount(tx, serviceModel.ClientID, serviceModel.ServiceTypeID)
		if err != nil && err != domain.ErrNotFound {
			return err
		}

		if clientServiceCount.ClientID != 0 {
			clientServiceCount.ServiceCount = clientServiceCount.ServiceCount + 1
			err := tx.
				Session(&gorm.Session{FullSaveAssociations: true}).
				Select("*").
				Where("client_id = ? AND service_type_id = ?", clientServiceCount.ClientID, clientServiceCount.ServiceTypeID).
				Updates(&clientServiceCount).Error
			if err != nil {
				return err
			}

			return nil
		}

		clientServiceCountModel := model.ClientServiceCount{
			ServiceTypeID: serviceModel.ServiceTypeID,
			ClientID:      serviceModel.ClientID,
			ServiceCount:  1,
		}

		err = db.Create(&clientServiceCountModel).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return domain.Service{}, err
	}

	return serviceModel.RepoToDomain(), nil
}

func (db *GormRepository) getClientServiceCount(tx *gorm.DB, cliendID uint, serviceType uint) (model.ClientServiceCount, error) {
	var clientServiceCount model.ClientServiceCount
	err := tx.Where("client_id = ? AND service_type_id = ?", cliendID, serviceType).Find(&clientServiceCount).Error
	if err != nil {
		return model.ClientServiceCount{}, err
	}
	if clientServiceCount.ClientID == 0 {
		return model.ClientServiceCount{}, domain.ErrNotFound
	}

	return clientServiceCount, nil
}
