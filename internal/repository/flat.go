package repository

import "github.com/go-pg/pg/v10"

type FlatRepository interface {
}

type flatRepository struct {
	db *pg.DB
}

func NewFlatRepository(db *pg.DB) FlatRepository {
	return &flatRepository{
		db: db,
	}
}
