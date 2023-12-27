package entity

type FlatGroupItem struct {
	FlatId  uint64 `gorm:"primaryKey"`
	GroupId uint64 `gorm:"primaryKey"`
	ItemId  uint64 `gorm:"primaryKey"`
	Status  string `gorm:"not null"`
}
