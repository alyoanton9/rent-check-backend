package repository

import (
	"github.com/go-pg/pg/v10"
	"rent-checklist/internal/models"
)

type FlatRepository interface {
	CreateFlat(flat *models.Flat) error
}

type flatRepository struct {
	db *pg.DB
}

func NewFlatRepository(db *pg.DB) FlatRepository {
	return &flatRepository{
		db: db,
	}
}

func (repo flatRepository) CreateFlat(flat *models.Flat) error {
	_, err := repo.db.Model(flat).Insert()
	if err != nil {
		return err
	}

	return nil
}
