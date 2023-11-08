package repository

import "github.com/go-pg/pg/v10"

type ItemRepository interface {
}

type itemRepository struct {
	db *pg.DB
}

func NewItemRepository(db *pg.DB) ItemRepository {
	return &itemRepository{db: db}
}
