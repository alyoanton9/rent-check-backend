package entity

type User struct {
	Id           uint64  `gorm:"primaryKey"`
	Login        string  `gorm:"index:,unique"`
	PasswordHash string  `gorm:"not null"`
	AuthToken    *string `gorm:"index:,unique"`
}
