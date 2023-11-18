package repository

import (
	"gorm.io/gorm"
	e "rent-checklist-backend/internal/error"
	"rent-checklist-backend/internal/model"
)

type FlatRepository interface {
	GetFlats() ([]model.Flat, error)
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

func (repo flatRepository) GetFlats() ([]model.Flat, error) {
	var flats []model.Flat
	err := repo.db.Find(&flats).Error

	if err != nil {
		return nil, err
	}

	return flats, nil
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
