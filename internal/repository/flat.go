package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	e "rent-checklist-backend/internal/error"
	"rent-checklist-backend/internal/model"
	"slices"
)

type FlatRepository interface {
	GetFlats(userId string) ([]model.Flat, error)
	CreateFlat(flat *model.Flat) error
	DeleteFlat(flatId uint64, userId string) error
	UpdateFlat(flat *model.Flat, userId string) error
}

type flatRepository struct {
	db *gorm.DB
}

func NewFlatRepository(db *gorm.DB) FlatRepository {
	return &flatRepository{
		db: db,
	}
}

func (repo flatRepository) GetFlats(userId string) ([]model.Flat, error) {
	var flatIds []uint64
	// TODO: check if raw "user_id = ?" can be replaced with
	// smth like model.UserFlat{UserId: userId}
	err := repo.db.Model(&[]model.UserFlat{}).Where("user_id = ?", userId).Pluck("flat_id", &flatIds).Error
	if err != nil {
		return nil, err
	}

	if len(flatIds) == 0 {
		return nil, nil
	}

	var flats []model.Flat
	err = repo.db.Where(flatIds).Find(&flats).Error
	if err != nil {
		return nil, err
	}

	return flats, nil
}

func (repo flatRepository) CreateFlat(flat *model.Flat) error {
	res := repo.db.Where(model.Flat{Address: flat.Address, OwnerId: flat.OwnerId}).FirstOrCreate(&flat)

	err := RaiseDbError(res, &e.KeyAlreadyExist{Msg: "unique", Field: "address"})
	if err != nil {
		return err
	}

	res = repo.db.Create(&model.UserFlat{UserId: flat.OwnerId, FlatId: flat.Id})

	err = RaiseDbError(res, &e.KeyAlreadyExist{Msg: "unique", Field: "flat_id,user_id"})
	if err != nil {
		return err
	}

	return err
}

func (repo flatRepository) DeleteFlat(flatId uint64, userId string) error {
	var ownerId string
	res := repo.db.Model(&model.Flat{}).Where(flatId).Pluck("owner_id", &ownerId)

	err := RaiseDbError(res, &e.KeyNotFound{Msg: "not-found", Field: "id"})
	if err != nil {
		return err
	}

	res = repo.db.Delete(&model.UserFlat{UserId: userId, FlatId: flatId})

	err = RaiseDbError(res, &e.ForbiddenAction{Msg: "has", Field: "flat_id"})
	if err != nil {
		return err
	}

	if ownerId == userId {
		err = repo.db.Delete(&model.Flat{}, flatId).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo flatRepository) UpdateFlat(flat *model.Flat, userId string) error {
	var userIds []string
	res := repo.db.Model(&model.UserFlat{}).Where(&model.UserFlat{FlatId: flat.Id}).Pluck("user_id", &userIds)

	err := RaiseDbError(res, &e.KeyNotFound{Msg: "not-found", Field: "id"})
	if err != nil {
		return err
	}

	if ok := slices.Contains(userIds, userId); !ok {
		return &e.ForbiddenAction{Msg: "has", Field: "flat_id"}
	}

	res = repo.db.Model(&flat).Clauses(clause.Returning{}).Updates(flat)

	return RaiseDbError(res, &e.KeyAlreadyExist{Msg: "unique", Field: "address"})
}
