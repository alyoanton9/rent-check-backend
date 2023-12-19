package entity

type FlatGroup struct {
	FlatId  uint64 `gorm:"primaryKey"`
	GroupId uint64 `gorm:"primaryKey"`
}
