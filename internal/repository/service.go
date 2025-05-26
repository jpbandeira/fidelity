package repository

// import (
// 	"github.com/google/uuid"
// 	"github.com/jp/fidelity/internal/domain"
// 	ferros "github.com/jp/fidelity/internal/pkg/errors"
// 	"github.com/jp/fidelity/internal/repository/model"
// 	"gorm.io/gorm"
// )











// func (db *GormRepository) GetClientServicesCount(cliendUUID string) ([]domain.ClientServiceTypeCount, error) {
// 	clientServicesCount := []model.ClientServiceTypeCount{}
// 	err := db.Preload("Client").Preload("ServiceType").Where("client_uuid = ?", cliendUUID).Find(&clientServicesCount).Error
// 	if err != nil {
// 		return []domain.ClientServiceTypeCount{}, err
// 	}

// 	var result []domain.ClientServiceTypeCount
// 	for _, csc := range clientServicesCount {
// 		result = append(result, csc.RepoToDomain())
// 	}

// 	return result, nil
// }

// func (db *GormRepository) ListServices(params []domain.Param) (domain.ServiceBatch, error) {
// 	var services []model.Service
// 	var q string
// 	var args []any

// 	for _, v := range params {
// 		q = q + v.Key + " = ?"
// 		args = append(args, v.Value)
// 	}

// 	err := db.Preload("Client").Preload("Attendant").Preload("ServiceType").Where(q, args...).Order("service_date DESC").Find(&services).Error
// 	if err != nil {
// 		return domain.ServiceBatch{}, err
// 	}

// 	client := services[0].Client
// 	attendant := services[0].Attendant
// 	return domain.ServiceBatch{
// 		Client:    client.RepoToDomain(),
// 		Attendant: attendant.RepoToDomain(),
// 		Items:     model.ServiceRepoToDomain(services),
// 	}, nil
// }
