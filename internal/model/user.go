package model

import (
	"rent-checklist-backend/internal/entity"
)

type User struct {
	Id           uint64
	Login        string
	PasswordHash string
	AuthToken    *string
}

func EntityToUser(user entity.User) User {
	return User{
		Id:           user.Id,
		Login:        user.Login,
		PasswordHash: user.PasswordHash,
		AuthToken:    user.AuthToken,
	}
}

func UserToEntity(user User) entity.User {
	return entity.User{
		Id:           user.Id,
		Login:        user.Login,
		PasswordHash: user.PasswordHash,
		AuthToken:    user.AuthToken,
	}
}
