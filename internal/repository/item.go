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

type ItemRepository interface {
	GetItems(userId string) ([]model.Item, error)
	CreateItem(item *model.Item) error
	UpdateItem(item *model.Item, userId string) error
	HideItem(itemId uint64, userId string) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (repo itemRepository) GetItems(userId string) ([]model.Item, error) {
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
	itemRecord := model.ItemToEntity(*item)

	err := repo.db.Create(&itemRecord).Error

	if errors.As(err, &gorm.ErrDuplicatedKey) {
		err = &e.KeyAlreadyExist{Msg: "unique", Field: "title"}
	}
	if err != nil {
		return err
	}

	*item = model.EntityToItem(itemRecord)

	return nil
}

func (repo itemRepository) UpdateItem(item *model.Item, userId string) error {
	err := checkUserHasItem(repo.db, item.Id, userId)
	if err != nil {
		return err
	}

	itemRecord := model.ItemToEntity(*item)

	err = repo.db.Model(&itemRecord).Clauses(clause.Returning{}).Updates(itemRecord).Error
	if errors.As(err, &gorm.ErrDuplicatedKey) {
		err = &e.KeyAlreadyExist{Msg: "unique", Field: "title"}
	}
	if err != nil {
		return err
	}

	*item = model.EntityToItem(itemRecord)

	return nil
}

func (repo itemRepository) HideItem(itemId uint64, userId string) error {
	err := checkUserHasItem(repo.db, itemId, userId)
	if err != nil {
		return err
	}

	err = repo.db.Model(&entity.Item{}).Where(itemId).Update("hide", true).Error
	return err
}

func checkUserHasItem(db *gorm.DB, itemId uint64, userId string) error {
	var itemUserId string
	err := db.Model(&entity.Item{}).Where(itemId).Pluck("user_id", &itemUserId).Error

	switch itemUserId {
	case "":
		err = &e.KeyNotFound{Msg: "not-found", Field: "itemId"}
	case userId:
		// ok
	default:
		err = &e.ForbiddenAction{Msg: "has", Field: "itemId"}
	}

	return err
}
