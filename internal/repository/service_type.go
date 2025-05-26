package repository

import (
	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/internal/repository/model"
	"gorm.io/gorm"
)

func (db *GormRepository) CreateServiceType(st domain.ServiceType) (domain.ServiceType, error) {
	var serviceTypeModel model.ServiceType

	err := db.Transaction(func(tx *gorm.DB) (err error) {
		serviceTypeModel = model.ServiceType{
			Name: st.Name,
		}

		return db.Create(&serviceTypeModel).Error
	})
	if err != nil {
		return domain.ServiceType{}, err
	}

	return serviceTypeModel.RepoToDomain(), nil
}

func (db *GormRepository) ListServiceTypes(params []domain.Param) ([]domain.ServiceType, error) {
	var serviceTypes []model.ServiceType
	var q string
	var args []interface{}

	for _, v := range params {
		q = q + v.Key + "=?"
		args = append(args, v.Value)
	}

	err := db.Where(q, args...).Find(&serviceTypes).Error
	if err != nil {
		return []domain.ServiceType{}, err
	}

	var result []domain.ServiceType
	for _, value := range serviceTypes {
		result = append(result, value.RepoToDomain())
	}

	return result, nil
}
