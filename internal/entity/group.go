package entity

type Group struct {
	Id     uint64 `gorm:"primaryKey"`
	UserId string `gorm:"not null"`
	Title  string `gorm:"not null"`
	Hide   bool   `gorm:"not null"`
}
