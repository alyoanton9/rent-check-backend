package repository

import (
	"errors"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"rent-checklist-backend/internal/entity"
	e "rent-checklist-backend/internal/error"
	"rent-checklist-backend/internal/model"
	"slices"
)

type FlatRepository interface {
	GetFlats(userId uint64) ([]model.Flat, error)
	CreateFlat(flat *model.Flat) error
	DeleteFlat(flatId uint64, userId uint64) error
	UpdateFlat(flat *model.Flat, userId uint64) error

	GetFlatsByGroupId(groupId, userId uint64) ([]model.Flat, error)
}

type flatRepository struct {
	db *gorm.DB
}

func NewFlatRepository(db *gorm.DB) FlatRepository {
	return &flatRepository{
		db: db,
	}
}

func (repo flatRepository) GetFlats(userId uint64) ([]model.Flat, error) {
	var flatIds []uint64
	err := repo.db.Model(&[]entity.UserFlat{}).Where(
		"user_id = ?", userId).Pluck("flat_id", &flatIds).Error
	if err != nil {
		return nil, err
	}

	if len(flatIds) == 0 {
		return nil, nil
	}

	var flatRecords []entity.Flat
	err = repo.db.Where(flatIds).Find(&flatRecords).Error
	if err != nil {
		return nil, err
	}

	flats := lo.Map(flatRecords, func(flat entity.Flat, _ int) model.Flat {
		return model.EntityToFlat(flat)
	})

	return flats, nil
}

func (repo flatRepository) CreateFlat(flat *model.Flat) error {
	flatRecord := model.FlatToEntity(*flat)

	err := repo.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&flatRecord).Error

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			err = &e.KeyAlreadyExist{Field: "address"}
		} else if errors.Is(err, gorm.ErrForeignKeyViolated) {
			err = &e.KeyNotFound{Field: "ownerId"}
		}
		if err != nil {
			return err
		}

		err = tx.Create(
			&entity.UserFlat{UserId: flatRecord.OwnerId, FlatId: flatRecord.Id}).Error

		return err
	})

	*flat = model.EntityToFlat(flatRecord)

	return err
}

func (repo flatRepository) DeleteFlat(flatId uint64, userId uint64) error {
	var ownerId uint64
	err := repo.db.Model(&entity.Flat{}).Where(flatId).Pluck("owner_id", &ownerId).Error

	if ownerId == 0 {
		err = &e.KeyNotFound{Field: "id"}
	}
	if err != nil {
		return err
	}

	err = repo.db.Transaction(func(tx *gorm.DB) error {
		res := repo.db.Delete(&entity.UserFlat{UserId: userId, FlatId: flatId})

		if res.RowsAffected == 0 {
			err = &e.NoAccess{Field: "userId,flatId"}
		}
		if err != nil {
			return err
		}

		if ownerId == userId {
			err = repo.db.Delete(&entity.Flat{}, flatId).Error
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func (repo flatRepository) UpdateFlat(flat *model.Flat, userId uint64) error {
	var userIds []uint64
	err := repo.db.Model(&entity.UserFlat{}).Where(
		&entity.UserFlat{FlatId: flat.Id}).Pluck("user_id", &userIds).Error

	if len(userIds) == 0 {
		err = &e.KeyNotFound{Field: "id"}
	}
	if err != nil {
		return err
	}

	if ok := slices.Contains(userIds, userId); !ok {
		return &e.NoAccess{Field: "userId,flatId"}
	}

	flatRecord := model.FlatToEntity(*flat)
	err = repo.db.Model(&flatRecord).Clauses(clause.Returning{}).Updates(flatRecord).Error

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		err = &e.KeyAlreadyExist{Field: "address"}
	}
	if err != nil {
		return err
	}

	*flat = model.EntityToFlat(flatRecord)

	return nil
}

func (repo flatRepository) GetFlatsByGroupId(groupId, userId uint64) ([]model.Flat, error) {
	err := checkUserHasGroup(repo.db, groupId, userId)
	if err != nil {
		return nil, err
	}

	var flatRecords []entity.Flat

	subQueryFlats := repo.db.Select("flat_id").Where("group_id=?", groupId).Table("flat_groups")
	subQueryUsers := repo.db.Select("flat_id").Where("user_id=?", userId).Table("user_flats")
	err = repo.db.Where("id IN (? INTERSECT ?)", subQueryFlats, subQueryUsers).Find(&flatRecords).Error
	if err != nil {
		return nil, err
	}

	flats := lo.Map(flatRecords, func(flat entity.Flat, _ int) model.Flat {
		return model.EntityToFlat(flat)
	})

	return flats, nil
}
