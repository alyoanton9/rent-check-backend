package entity

type User struct {
	Id        string `gorm:"primaryKey"`
	AuthToken string `gorm:"index:,unique"`
}
