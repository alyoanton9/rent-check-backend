package entity

type Item struct {
	Id          uint64 `gorm:"primaryKey"`
	UserId      uint64 `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string
	Hide        bool `gorm:"not null"`
}
