package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jp/fidelity/internal/domain"
	ferros "github.com/jp/fidelity/internal/pkg/errors"
	"github.com/jp/fidelity/internal/repository/model"
)

func (db *GormRepository) CreateAttendant(attendant domain.Attendant) (domain.Attendant, error) {
	attendantModel := model.Attendant{
		UUID:  uuid.NewString(),
		Name:  attendant.Name,
		Email: attendant.Email,
		Phone: attendant.Phone,
	}

	err := db.Create(&attendantModel).Error
	if err != nil {
		return domain.Attendant{}, err
	}

	return attendantModel.RepoToDomain(), nil
}

func (db *GormRepository) UpdateAttendant(attendant domain.Attendant) (domain.Attendant, error) {
	var newAttendant model.Attendant

	err := db.Transaction(func(tx *gorm.DB) (err error) {
		oldAttendant, err := db.getAttendant(tx, attendant.ID)
		if err != nil {
			return err
		}

		oldAttendant.Name = attendant.Name
		oldAttendant.Email = attendant.Email
		oldAttendant.Phone = attendant.Phone
		oldAttendant.UpdatedAt = time.Now()

		t := tx.
			Session(&gorm.Session{FullSaveAssociations: true}).
			Select("*").
			Updates(&oldAttendant)
		if t.Error != nil {
			return t.Error
		}
		if t.RowsAffected == 0 {
			return ferros.ErrNotFound
		}
		newAttendant = oldAttendant
		return nil
	})
	if err != nil {
		return domain.Attendant{}, err
	}

	return newAttendant.RepoToDomain(), nil
}

func (db *GormRepository) GetAttendant(uuid string) (domain.Attendant, error) {
	var attendant model.Attendant

	err := db.Transaction(func(tx *gorm.DB) (err error) {
		attendant, err = db.getAttendant(tx, uuid)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return domain.Attendant{}, err
	}

	if attendant.ID == 0 {
		return domain.Attendant{}, ferros.ErrNotFound
	}

	return attendant.RepoToDomain(), nil
}

func (db *GormRepository) getAttendant(tx *gorm.DB, uuid string) (model.Attendant, error) {
	var attendant model.Attendant

	err := tx.Where("uuid = ?", uuid).Find(&attendant).Error
	if err != nil {
		return model.Attendant{}, err
	}

	return attendant, nil
}

func (db *GormRepository) ListAttendants(params []domain.Param) ([]domain.Attendant, error) {
	var attendants []model.Attendant
	var q string
	var args []interface{}

	for _, v := range params {
		q = q + v.Key + "=?"
		args = append(args, v.Value)
	}

	err := db.Where(q, args...).Find(&attendants).Error
	if err != nil {
		return []domain.Attendant{}, err
	}

	var result []domain.Attendant
	for _, value := range attendants {
		result = append(result, value.RepoToDomain())
	}

	return result, nil
}

func (db *GormRepository) DeleteAttendant(uuid string) error {
	tx := db.Unscoped().Where(&model.Attendant{UUID: uuid}).Delete(&model.Attendant{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return ferros.ErrNotFound
	}

	return nil
}
