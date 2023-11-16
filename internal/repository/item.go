package repository

import (
	"gorm.io/gorm"
)

type ItemRepository interface {
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}
