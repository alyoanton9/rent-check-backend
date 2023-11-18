package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rent-checklist-backend/internal/repository"
)

type Handler interface {
	GetFlats(ctx echo.Context) error
	CreateFlat(ctx echo.Context) error

	HomePage(c echo.Context) error
}

type handler struct {
	flatRepository repository.FlatRepository
	itemRepository repository.ItemRepository
}

func NewHandler(flat repository.FlatRepository, item repository.ItemRepository) Handler {
	return &handler{flatRepository: flat, itemRepository: item}
}

func (h handler) HomePage(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
