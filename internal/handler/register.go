package handler

import (
	"github.com/labstack/echo/v4"
	"log"
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/model"
)

func (h handler) Register(ctx echo.Context) error {
	var userRequest *dto.UserRequest
	err := ParseBody(ctx, &userRequest, "register user request")
	if err != nil {
		return err
	}

	passwordHash, err := h.hasher.HashPassword(userRequest.Password)
	if err != nil {
		log.Printf("error hashing password: %s", err.Error())
		return echo.ErrInternalServerError
	}

	user := model.User{
		Login:        userRequest.Login,
		PasswordHash: passwordHash,
	}

	err = h.userRepository.CreateUser(&user)

	if err != nil {
		return HandleDbError(ctx, err, "error registering user")
	}

	return nil
}
