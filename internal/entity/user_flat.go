package entity

type UserFlat struct {
	UserId uint64 `gorm:"primaryKey"`
	FlatId uint64 `gorm:"primaryKey"`
}
