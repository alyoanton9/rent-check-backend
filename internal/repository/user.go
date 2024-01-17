package repository

import (
	"errors"
	"gorm.io/gorm"
	"rent-checklist-backend/internal/entity"
	e "rent-checklist-backend/internal/error"
	"rent-checklist-backend/internal/model"
)

type UserRepository interface {
	GetUserById(id uint64) (*model.User, error)
	GetUserByAuthToken(authToken string) (*model.User, error)
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{db: db}
}

func (repo userRepository) GetUserById(id uint64) (*model.User, error) {
	userRecord := entity.User{Id: id}

	err := repo.db.First(&userRecord).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = &e.KeyNotFound{Field: "id"}
	}
	if err != nil {
		return nil, err
	}

	user := model.EntityToUser(userRecord)

	return &user, nil
}

func (repo userRepository) GetUserByAuthToken(authToken string) (*model.User, error) {
	var userRecord entity.User
	err := repo.db.Where(&entity.User{AuthToken: &authToken}).First(&userRecord).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = &e.KeyNotFound{Field: "authToken"}
	}
	if err != nil {
		return nil, err
	}

	user := model.EntityToUser(userRecord)

	return &user, nil

}

func (repo userRepository) CreateUser(user *model.User) error {
	userRecord := model.UserToEntity(*user)
	err := repo.db.Where(&entity.User{Login: userRecord.Login}).First(&entity.User{}).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &e.KeyAlreadyExist{Field: "login"}
	}

	err = repo.db.Create(&userRecord).Error

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		err = &e.KeyAlreadyExist{Field: "id"}
	}
	if err != nil {
		return err
	}

	*user = model.EntityToUser(userRecord)

	return err
}
