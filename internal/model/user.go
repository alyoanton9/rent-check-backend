package model

type User struct {
	Id        string `json:"id" gorm:"primaryKey"`
	AuthToken string `json:"auth_token"`
}
