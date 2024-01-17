package entity

type Flat struct {
	Id          uint64 `gorm:"primaryKey"`
	Title       string
	Description string
	Address     string `gorm:"not null"`
	OwnerId     uint64 `gorm:"not null"`
}
