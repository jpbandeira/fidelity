package repository

import (
	"github.com/google/uuid"
	"github.com/jp/fidelity/internal/domain"
	ferros "github.com/jp/fidelity/internal/pkg/errors"
	"github.com/jp/fidelity/internal/repository/model"
	"gorm.io/gorm"
)

func (db *GormRepository) CreateServiceBatch(serviceBatch domain.ServiceBatch) (domain.ServiceBatch, error) {
	var serviceModelList = make([]model.Service, 0)
	err := db.Transaction(func(tx *gorm.DB) (err error) {
		for _, s := range serviceBatch.Items {
			serviceType, err := db.getServiceType(tx, s.ServiceType)
			if err != nil {
				return err
			}

			serviceModel := model.Service{
				UUID:          uuid.NewString(),
				ClientUUID:    serviceBatch.Client.ID,
				Client:        model.ClientDomainToRepo(serviceBatch.Client),
				AttedantUUID:  serviceBatch.Attendant.ID,
				Attendant:     model.AttendantDomainToRepo(serviceBatch.Attendant),
				Price:         s.Price,
				ServiceTypeID: serviceType.ID,
				ServiceType:   serviceType,
				PaymentType:   domain.ToPaymentType(s.PaymentType),
				Description:   s.Description,
				ServiceDate:   s.ServiceDate,
			}

			serviceModelList = append(serviceModelList, serviceModel)
		}

		err = db.Create(&serviceModelList).Error
		if err != nil {
			return err
		}

		for _, s := range serviceModelList {
			clientServiceCount, cscErr := db.getClientServiceCount(tx, s.Client.UUID, s.ServiceTypeID)
			if cscErr != nil && cscErr != ferros.ErrNotFound {
				return cscErr
			}

			if clientServiceCount.ClientUUID != "" {
				return db.updateClientServiceCount(tx, clientServiceCount)
			}

			err := db.createClientServiceCount(tx, s.ServiceTypeID, serviceBatch.Client.ID)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return domain.ServiceBatch{}, err
	}

	client := serviceModelList[0].Client
	attendant := serviceModelList[0].Attendant
	return domain.ServiceBatch{
		Client:    client.RepoToDomain(),
		Attendant: attendant.RepoToDomain(),
		Items:     model.ServiceRepoToDomain(serviceModelList),
	}, nil
}

func (db *GormRepository) getServiceType(tx *gorm.DB, serviceType string) (model.ServiceType, error) {
	var serviceTypeModel = model.ServiceType{}
	t := tx.Where("description = ?", serviceType).Find(&serviceTypeModel)
	if t.Error != nil {
		return model.ServiceType{}, t.Error
	}
	if serviceTypeModel.ID == 0 {
		return model.ServiceType{}, ferros.ErrNotFound
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
		return model.ClientServiceTypeCount{}, ferros.ErrNotFound
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

func (db *GormRepository) ListServices(params []domain.Param) (domain.ServiceBatch, error) {
	var services []model.Service
	var q string
	var args []any

	for _, v := range params {
		q = q + v.Key + " = ?"
		args = append(args, v.Value)
	}

	err := db.Preload("Client").Preload("Attendant").Preload("ServiceType").Where(q, args...).Order("service_date DESC").Find(&services).Error
	if err != nil {
		return domain.ServiceBatch{}, err
	}

	client := services[0].Client
	attendant := services[0].Attendant
	return domain.ServiceBatch{
		Client:    client.RepoToDomain(),
		Attendant: attendant.RepoToDomain(),
		Items:     model.ServiceRepoToDomain(services),
	}, nil
}
