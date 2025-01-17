package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jp/fidelity/internal/domain"
	"github.com/jp/fidelity/internal/repository/model"
)

func (db *GormRepository) CreateUser(user domain.User) (domain.User, error) {
	userModel := model.User{
		UUID:  uuid.NewString(),
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}

	err := db.Create(&userModel).Error
	if err != nil {
		return domain.User{}, err
	}

	return userModel.RepoToDomain(), nil
}

func (db *GormRepository) UpdateUser(user domain.User) (domain.User, error) {
	var newUser model.User

	err := db.Transaction(func(tx *gorm.DB) (err error) {
		oldUser, err := db.getUser(tx, user.ID)
		if err != nil {
			return err
		}

		oldUser.Name = user.Name
		oldUser.Email = user.Email
		oldUser.Phone = user.Phone
		oldUser.UpdatedAt = time.Now()

		t := tx.
			Session(&gorm.Session{FullSaveAssociations: true}).
			Select("*").
			Updates(&oldUser)
		if t.Error != nil {
			return t.Error
		}
		if t.RowsAffected == 0 {
			return domain.ErrNotFound
		}
		newUser = oldUser
		return nil
	})
	if err != nil {
		return domain.User{}, err
	}

	return newUser.RepoToDomain(), nil
}

func (db *GormRepository) GetUser(uuid string) (domain.User, error) {
	var user model.User

	err := db.Transaction(func(tx *gorm.DB) (err error) {
		user, err = db.getUser(tx, uuid)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return domain.User{}, err
	}

	if user.ID == 0 {
		return domain.User{}, domain.ErrNotFound
	}

	return user.RepoToDomain(), nil
}

func (db *GormRepository) getUser(tx *gorm.DB, uuid string) (model.User, error) {
	var user model.User

	err := tx.Where("uuid = ?", uuid).Find(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (db *GormRepository) ListUsers(params []domain.Param) ([]domain.User, error) {
	var users []model.User
	var q string
	var args []interface{}

	for _, v := range params {
		q = q + v.Key + "=?"
		args = append(args, v.Value)
	}

	err := db.Where(q, args...).Find(&users).Error
	if err != nil {
		return []domain.User{}, err
	}

	var result []domain.User
	for _, value := range users {
		result = append(result, value.RepoToDomain())
	}

	return result, nil
}

func (db *GormRepository) DeleteUser(uuid string) error {
	tx := db.Unscoped().Where(&model.User{UUID: uuid}).Delete(&model.User{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}
