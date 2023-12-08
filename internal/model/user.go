package model

import "rent-checklist-backend/internal/entity"

type User struct {
	Id        string
	AuthToken string
}

func EntityToUser(user entity.User) User {
	return User{
		Id:        user.Id,
		AuthToken: user.AuthToken,
	}
}

func UserToEntity(user User) entity.User {
	return entity.User{
		Id:        user.Id,
		AuthToken: user.AuthToken,
	}
}
