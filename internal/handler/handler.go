package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rent-checklist/internal/repository"
)

type Handler interface {
	GetFlats(c echo.Context) error
	CreateFlat(c echo.Context) error

	HomePage(c echo.Context) error
}

type handler struct {
	flatRepository repository.FlatRepository
	itemRepository repository.ItemRepository
}

func NewHandler(flat repository.FlatRepository, item repository.ItemRepository) Handler {
	return &handler{flatRepository: flat, itemRepository: item}
}

func (h handler) HomePage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
