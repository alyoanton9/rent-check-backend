package repository

import (
	"errors"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"rent-checklist-backend/internal/entity"
	e "rent-checklist-backend/internal/error"
	"rent-checklist-backend/internal/model"
)

type GroupRepository interface {
	GetGroups(userId string, ids []uint64) ([]model.Group, error)
	CreateGroup(group *model.Group) error
	UpdateGroup(group *model.Group, userId string) error
	AddGroupToFlat(groupId, flatId uint64, userId string) error
	DeleteGroupFromFlat(groupId uint64, flatId uint64, userId string) error
	HideGroup(groupId uint64, userId string) error
}

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return groupRepository{db: db}
}

func (repo groupRepository) GetGroups(userId string, ids []uint64) ([]model.Group, error) {
	var groupRecords []entity.Group

	err := repo.db.Model(&[]entity.Group{}).Where(ids).Where(
		"user_id = ? AND hide = false", userId).Find(&groupRecords).Error
	if err != nil {
		return nil, err
	}

	// TODO what is the correct way to return empty list -- nil or {}?
	if len(groupRecords) == 0 {
		return nil, nil
	}

	groups := lo.Map(groupRecords, func(group entity.Group, _ int) model.Group {
		return model.EntityToGroup(group)
	})

	return groups, nil
}

func (repo groupRepository) CreateGroup(group *model.Group) error {
	groupRecord := model.GroupToEntity(*group)

	err := repo.db.Create(&groupRecord).Error

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		err = &e.KeyAlreadyExist{Field: "title"}
	} else if errors.Is(err, gorm.ErrForeignKeyViolated) {
		err = &e.KeyNotFound{Field: "userId"}
	}
	if err != nil {
		return err
	}

	*group = model.EntityToGroup(groupRecord)

	return nil
}

func (repo groupRepository) UpdateGroup(group *model.Group, userId string) error {
	err := checkUserHasGroup(repo.db, group.Id, userId)
	if err != nil {
		return err
	}

	groupRecord := model.GroupToEntity(*group)

	err = repo.db.Model(&groupRecord).Clauses(clause.Returning{}).Updates(groupRecord).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		err = &e.KeyAlreadyExist{Field: "title"}
	}
	if err != nil {
		return err
	}

	*group = model.EntityToGroup(groupRecord)

	return nil
}

func (repo groupRepository) AddGroupToFlat(groupId, flatId uint64, userId string) error {
	err := checkUserHasFlat(repo.db, userId, flatId)
	if err != nil {
		return err
	}

	err = checkUserHasGroup(repo.db, groupId, userId)
	if err != nil {
		return err
	}

	err = repo.db.Create(&entity.FlatGroup{FlatId: flatId, GroupId: groupId}).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		err = &e.KeyAlreadyExist{Field: "groupId"}
	}

	return err
}

func (repo groupRepository) DeleteGroupFromFlat(groupId, flatId uint64, userId string) error {
	err := checkUserHasGroup(repo.db, groupId, userId)
	if err != nil {
		return err
	}

	err = repo.db.Transaction(func(tx *gorm.DB) error {
		var err error

		res := tx.Delete(&entity.FlatGroup{FlatId: flatId, GroupId: groupId})
		if res.RowsAffected == 0 {
			err = &e.NoAccess{Field: "flatId,groupId"}
		}
		if err != nil {
			return err
		}

		err = tx.Where("flat_id = ? AND group_id = ?", flatId, groupId).Delete(&entity.FlatGroupItem{}).Error

		return err
	})

	return err
}

func (repo groupRepository) HideGroup(groupId uint64, userId string) error {
	err := checkUserHasGroup(repo.db, groupId, userId)
	if err != nil {
		return err
	}

	err = repo.db.Model(&entity.Group{}).Where(groupId).Update("hide", true).Error
	return err
}

func checkUserHasGroup(db *gorm.DB, groupId uint64, userId string) error {
	groupRecord := entity.Group{Id: groupId}
	err := db.First(&groupRecord).Error

	switch {
	case groupRecord.UserId == "" || groupRecord.Hide:
		err = &e.KeyNotFound{Field: "groupId"}
	case groupRecord.UserId == userId:
		// ok
	default:
		err = &e.NoAccess{Field: "userId,groupId"}
	}

	return err
}

func checkUserHasFlat(db *gorm.DB, userId string, flatId uint64) error {
	var userFlat entity.UserFlat
	err := db.Where(entity.UserFlat{UserId: userId, FlatId: flatId}).First(&userFlat).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = &e.NoAccess{Field: "userId,flatId"}
	}

	return err
}
