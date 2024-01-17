package repository

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"rent-checklist-backend/internal/entity"
	e "rent-checklist-backend/internal/error"
	"rent-checklist-backend/internal/model"
)

type ItemRepository interface {
	GetItems(userId uint64) ([]model.Item, error)
	CreateItem(item *model.Item) error
	UpdateItem(item *model.Item, userId uint64) error
	HideItem(itemId uint64, userId uint64) error
	AddItemToGroup(flatId, groupId, itemId uint64, userId uint64) error
	DeleteItemFromGroup(flatId, groupId, itemId uint64, userId uint64) error
	GetFlatItems(flatId uint64, userId uint64) ([]model.GroupItems, error)
	UpdateItemStatus(flatId, groupId, itemId uint64, status model.Status, userId uint64) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (repo itemRepository) GetItems(userId uint64) ([]model.Item, error) {
	var itemRecords []entity.Item
	err := repo.db.Model(&[]entity.Item{}).Where(
		"user_id = ? AND hide = false", userId).Find(&itemRecords).Error
	if err != nil {
		return nil, err
	}

	if len(itemRecords) == 0 {
		return nil, nil
	}

	items := lo.Map(itemRecords, func(item entity.Item, _ int) model.Item {
		return model.EntityToItem(item)
	})

	return items, nil
}

func (repo itemRepository) CreateItem(item *model.Item) error {
	existingItemRecord := entity.Item{}

	err := repo.db.Where("title = ?", item.Title).First(&existingItemRecord).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		itemRecord := model.ItemToEntity(*item)
		err = repo.db.Create(&itemRecord).Error

		*item = model.EntityToItem(itemRecord)
	} else {
		if existingItemRecord.Hide {
			existingItemRecord.Hide = false
			err = repo.db.Save(&existingItemRecord).Error

			*item = model.EntityToItem(existingItemRecord)
		} else {
			err = &e.KeyAlreadyExist{Field: "title"}
		}
	}

	return err
}

func (repo itemRepository) UpdateItem(item *model.Item, userId uint64) error {
	err := checkUserHasItem(repo.db, item.Id, userId)
	if err != nil {
		return err
	}

	itemRecord := model.ItemToEntity(*item)

	err = repo.db.Model(&itemRecord).Clauses(clause.Returning{}).Updates(itemRecord).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		err = &e.KeyAlreadyExist{Field: "title"}
	}
	if err != nil {
		return err
	}

	*item = model.EntityToItem(itemRecord)

	return nil
}

func (repo itemRepository) HideItem(itemId uint64, userId uint64) error {
	err := checkUserHasItem(repo.db, itemId, userId)
	if err != nil {
		return err
	}

	err = repo.db.Model(&entity.Item{}).Where(itemId).Update("hide", true).Error
	return err
}

func (repo itemRepository) AddItemToGroup(flatId, groupId, itemId uint64, userId uint64) error {
	err := checkUserHasFlat(repo.db, userId, flatId)
	if err != nil {
		return err
	}

	err = checkFlatHasGroup(repo.db, flatId, groupId)
	if err != nil {
		return err
	}

	err = checkUserHasItem(repo.db, itemId, userId)
	if err != nil {
		return err
	}

	var initialStatus model.Status = 0
	err = repo.db.Create(&entity.FlatGroupItem{
		FlatId: flatId, GroupId: groupId, ItemId: itemId, Status: initialStatus.String(),
	}).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		err = &e.KeyAlreadyExist{Field: "flatId,groupId,itemId"}
	}

	return err
}

func (repo itemRepository) DeleteItemFromGroup(flatId, groupId, itemId uint64, userId uint64) error {
	err := checkUserHasFlat(repo.db, userId, flatId)
	if err != nil {
		return err
	}

	err = checkFlatHasGroup(repo.db, flatId, groupId)
	if err != nil {
		return err
	}

	err = checkUserHasItem(repo.db, itemId, userId)
	if err != nil {
		return err
	}

	res := repo.db.Delete(&entity.FlatGroupItem{FlatId: flatId, GroupId: groupId, ItemId: itemId})
	if res.RowsAffected == 0 {
		err = &e.NoAccess{Field: "flatId,groupId,itemId"} // TODO err msg needs refactoring
	}

	return err
}

func (repo itemRepository) GetFlatItems(flatId uint64, userId uint64) ([]model.GroupItems, error) {
	err := checkUserHasFlat(repo.db, userId, flatId)
	if err != nil {
		return nil, err
	}

	groupItemsRecords := make([]entity.GroupItem, 0)

	query := fmt.Sprintf(`
		select items.id, title, description, user_id, hide, status, flat_groups.group_id
		from flat_groups
		left join flat_group_items on flat_groups.group_id=flat_group_items.group_id
		left join items on items.id = flat_group_items.item_id
		where flat_groups.flat_id=%d`, flatId)

	err = repo.db.Raw(query).Scan(&groupItemsRecords).Error

	if err != nil {
		return nil, err
	}

	return model.EntitiesToGroupItems(groupItemsRecords), nil
}

func (repo itemRepository) UpdateItemStatus(flatId, groupId, itemId uint64, status model.Status, userId uint64) error {
	err := checkUserHasFlat(repo.db, userId, flatId)
	if err != nil {
		return err
	}

	err = checkFlatHasGroup(repo.db, flatId, groupId)
	if err != nil {
		return err
	}

	err = checkUserHasItem(repo.db, itemId, userId)
	if err != nil {
		return err
	}

	err = repo.db.Model(&entity.FlatGroupItem{
		FlatId: flatId, GroupId: groupId, ItemId: itemId,
	}).Update("status", status).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = &e.KeyNotFound{Field: "flatId,groupId,itemId"}
	}

	return err
}

func checkUserHasItem(db *gorm.DB, itemId uint64, userId uint64) error {
	itemRecord := entity.Item{Id: itemId}
	err := db.First(&itemRecord).Error

	switch {
	case itemRecord.UserId == 0 || itemRecord.Hide:
		err = &e.KeyNotFound{Field: "itemId"}
	case itemRecord.UserId == userId:
		// ok
	default:
		err = &e.NoAccess{Field: "userId,itemId"}
	}

	return err
}

func checkFlatHasGroup(db *gorm.DB, flatId, groupId uint64) error {
	var flatGroup entity.FlatGroup
	err := db.Where(entity.FlatGroup{FlatId: flatId, GroupId: groupId}).First(&flatGroup).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = &e.NoAccess{Field: "flatId,groupId"}
	}

	return err
}
