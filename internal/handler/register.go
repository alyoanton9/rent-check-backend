package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"rent-checklist-backend/internal/model"
	"strings"
)

func (h handler) RegisterUser(ctx echo.Context) error {
	authHeader := ctx.Request().Header.Get("Authorization")

	// no auth header provided, TODO: is verbose error response with msg needed?
	if authHeader == "" {
		return echo.ErrBadRequest
	}

	authToken, err := getAuthToken(authHeader, h.authService.GetAuthScheme())
	if err != nil {
		return echo.ErrBadRequest
	}

	userId, err := h.authService.AuthenticateUser(authToken, "")
	if err != nil {
		return echo.ErrInternalServerError
	}

	user := model.User{AuthToken: authToken, Id: userId}

	err = h.userRepository.CreateUser(&user)
	if err != nil {
		return HandleDbError(ctx, err, "error registering user")
	}

	return nil
}

// TODO perhaps move to auth service
func getAuthToken(authHeader string, authScheme string) (string, error) {
	prefix := fmt.Sprintf("%s ", authScheme)
	authToken := strings.TrimPrefix(authHeader, prefix)

	if authToken == authHeader {
		return "", fmt.Errorf("incorrect format of auth header")
	}

	return authToken, nil
}
