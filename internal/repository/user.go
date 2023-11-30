package repository

import (
	"gorm.io/gorm"
	e "rent-checklist-backend/internal/error"
	"rent-checklist-backend/internal/model"
)

type UserRepository interface {
	GetUserById(id string) (*model.User, error)
	GetUserByAuthToken(authToken string) (*model.User, error)
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{db: db}
}

func (repo userRepository) GetUserById(id string) (*model.User, error) {
	user := model.User{Id: id}
	res := repo.db.First(&user)

	err := RaiseDbError(res, &e.KeyNotFound{Msg: "not-found", Field: "id"})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo userRepository) GetUserByAuthToken(authToken string) (*model.User, error) {
	var user model.User
	res := repo.db.Where(&model.User{AuthToken: authToken}).First(&user)

	err := RaiseDbError(res, &e.KeyNotFound{Msg: "not-found", Field: "auth_token"})
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (repo userRepository) CreateUser(user *model.User) error {
	var existingUser model.User
	res := repo.db.Where(&model.User{AuthToken: user.AuthToken}).First(&existingUser)

	err := res.Error
	if res.RowsAffected > 0 {
		err = &e.KeyAlreadyExist{Msg: "unique", Field: "auth_token"}
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	res = repo.db.Where(user.Id).FirstOrCreate(&user)

	err = RaiseDbError(res, &e.KeyAlreadyExist{Msg: "unique", Field: "id"})

	return err
}
