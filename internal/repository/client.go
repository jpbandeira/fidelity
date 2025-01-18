package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/internal/repository/model"
)

func (db *GormRepository) CreateClient(client domain.Client) (domain.Client, error) {
	clientModel := model.Client{
		UUID:  uuid.NewString(),
		Name:  client.Name,
		Email: client.Email,
		Phone: client.Phone,
	}

	err := db.Create(&clientModel).Error
	if err != nil {
		return domain.Client{}, err
	}

	return clientModel.RepoToDomain(), nil
}

func (db *GormRepository) UpdateClient(client domain.Client) (domain.Client, error) {
	var newClient model.Client

	err := db.Transaction(func(tx *gorm.DB) (err error) {
		oldClient, err := db.getClient(tx, client.ID)
		if err != nil {
			return err
		}

		oldClient.Name = client.Name
		oldClient.Email = client.Email
		oldClient.Phone = client.Phone
		oldClient.UpdatedAt = time.Now()

		t := tx.
			Session(&gorm.Session{FullSaveAssociations: true}).
			Select("*").
			Updates(&oldClient)
		if t.Error != nil {
			return t.Error
		}
		if t.RowsAffected == 0 {
			return domain.ErrNotFound
		}
		newClient = oldClient
		return nil
	})
	if err != nil {
		return domain.Client{}, err
	}

	return newClient.RepoToDomain(), nil
}

func (db *GormRepository) GetClient(uuid string) (domain.Client, error) {
	var client model.Client

	err := db.Transaction(func(tx *gorm.DB) (err error) {
		client, err = db.getClient(tx, uuid)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return domain.Client{}, err
	}

	if client.ID == 0 {
		return domain.Client{}, domain.ErrNotFound
	}

	return client.RepoToDomain(), nil
}

func (db *GormRepository) getClient(tx *gorm.DB, uuid string) (model.Client, error) {
	var client model.Client

	err := tx.Where("uuid = ?", uuid).Find(&client).Error
	if err != nil {
		return model.Client{}, err
	}

	return client, nil
}

func (db *GormRepository) ListClients(params []domain.Param) ([]domain.Client, error) {
	var clients []model.Client
	var q string
	var args []interface{}

	for _, v := range params {
		q = q + v.Key + "=?"
		args = append(args, v.Value)
	}

	err := db.Where(q, args...).Find(&clients).Error
	if err != nil {
		return []domain.Client{}, err
	}

	var result []domain.Client
	for _, value := range clients {
		result = append(result, value.RepoToDomain())
	}

	return result, nil
}

func (db *GormRepository) DeleteClient(uuid string) error {
	tx := db.Unscoped().Where(&model.Client{UUID: uuid}).Delete(&model.Client{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}
