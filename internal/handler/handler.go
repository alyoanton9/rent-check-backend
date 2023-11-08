package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rent-checklist/internal/repository"
)

type Handler interface {
	HomePage(c echo.Context) error
}

type handler struct {
	flat repository.FlatRepository
	item repository.ItemRepository
}

func NewHandler(flat repository.FlatRepository, item repository.ItemRepository) Handler {
	return &handler{flat: flat, item: item}
}

func (h handler) HomePage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
