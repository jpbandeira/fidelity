package repository

import (
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/internal/repository/model"
	"gorm.io/gorm"
)

func (db *GormRepository) CreateService(service model.Service) (domain.Service, error) {
	serviceModel := model.Service{}

	err := db.Transaction(func(tx *gorm.DB) error {
		err := db.Create(&serviceModel).Error
		if err != nil {
			return err
		}

		// // clientServiceCount, err := getServiceCountByClientIDAndSerciceType()
		// caso o erro seja not found, deve pular o IF abaixo
		// caso não seja not found mas é algum erro de execução, o erro deve ser retornado
		// no caso de não ter nenhum erro é porque encontrou e deve ser atualizado
		// if err != nil && err != "not found" {
		// 	return err
		// }

		// if err != "not found" {
		// 	clientServiceCount.ServiceCount++
		// 	err := tx.
		// 		Session(&gorm.Session{FullSaveAssociations: true}).
		// 		Select("*").
		// 		Updates(&oldUser).Error
		// 	if err != nil {
		// 		return t.Error
		// 	}

		// return nil
		// }

		clientServiceCountModel := model.ClientServiceCount{
			ServiceType:  serviceModel.ServiceType,
			ClientID:     serviceModel.ClientID,
			ServiceCount: 1,
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
