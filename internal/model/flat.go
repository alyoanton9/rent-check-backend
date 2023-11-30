package model

type Flat struct {
	Id          uint64 `gorm:"primaryKey"`
	Title       string
	Description string
	Address     string
	OwnerId     string
}
