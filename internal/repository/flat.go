package repository

import (
	"gorm.io/gorm"
	e "rent-checklist/internal/error"
	"rent-checklist/internal/model"
)

type FlatRepository interface {
	CreateFlat(flat *model.Flat) error
}

type flatRepository struct {
	db *gorm.DB
}

func NewFlatRepository(db *gorm.DB) FlatRepository {
	return &flatRepository{
		db: db,
	}
}

func (repo flatRepository) CreateFlat(flat *model.Flat) error {
	res := repo.db.Where(model.Flat{Address: flat.Address}).FirstOrCreate(flat)

	err := res.Error
	if res.RowsAffected == 0 {
		err = &e.KeyAlreadyExist{Msg: "unique", Field: "address"}
	}

	if err != nil {
		return err
	}

	return nil
}
