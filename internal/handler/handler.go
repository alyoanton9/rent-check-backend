package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rent-checklist-backend/internal/repository"
	"rent-checklist-backend/internal/service"
)

type Handler interface {
	GetFlats(ctx echo.Context) error
	CreateFlat(ctx echo.Context) error
	DeleteFlat(ctx echo.Context) error
	UpdateFlat(ctx echo.Context) error

	HomePage(ctx echo.Context) error
	RegisterUser(ctx echo.Context) error
}

type handler struct {
	userRepository repository.UserRepository
	flatRepository repository.FlatRepository
	itemRepository repository.ItemRepository
	authService    service.AuthService
}

func NewHandler(user repository.UserRepository, flat repository.FlatRepository, item repository.ItemRepository, auth service.AuthService) Handler {
	return &handler{userRepository: user, flatRepository: flat, itemRepository: item, authService: auth}
}

func (h handler) HomePage(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
