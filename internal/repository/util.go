package repository

import (
	"gorm.io/gorm"
)

func RaiseDbError(res *gorm.DB, dbErr error) error {
	err := res.Error
	if res.RowsAffected == 0 {
		err = dbErr
	}

	return err
}
