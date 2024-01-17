package handler

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"rent-checklist-backend/internal/dto"
)

func (h handler) Login(ctx echo.Context) error {
	var userRequest *dto.UserRequest
	err := ParseBody(ctx, &userRequest, "register user request")
	if err != nil {
		return err
	}

	user, err := h.userRepository.GetUserByLogin(userRequest.Login)

	passwordsMatch := h.hasher.PasswordsMatch(userRequest.Password, user.PasswordHash)
	if !passwordsMatch {
		log.Printf("passwords don't match")
		return echo.ErrUnauthorized
	}

	authToken, err := h.authService.CreateToken(userRequest.Login)
	if err != nil {
		log.Printf("error generating auth token: %s", err.Error())
		return echo.ErrInternalServerError
	}

	user.AuthToken = &authToken
	err = h.userRepository.UpdateUser(user)
	if err != nil {
		return HandleDbError(ctx, err, "error updating user's auth token")
	}

	tokenResponse := dto.TokenResponse{Token: authToken}

	return ctx.JSON(http.StatusOK, tokenResponse)
}
