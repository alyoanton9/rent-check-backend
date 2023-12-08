package entity

type UserFlat struct {
	UserId string `gorm:"primaryKey"`
	FlatId uint64 `gorm:"primaryKey"`
}
