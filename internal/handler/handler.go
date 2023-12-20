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
	UpdateFlat(ctx echo.Context) error
	DeleteFlat(ctx echo.Context) error

	GetGroups(ctx echo.Context) error
	CreateGroup(ctx echo.Context) error
	UpdateGroup(ctx echo.Context) error
	AddGroupToFlat(ctx echo.Context) error
	DeleteGroupFromFlat(ctx echo.Context) error
	HideGroup(ctx echo.Context) error

	GetItems(ctx echo.Context) error
	CreateItem(ctx echo.Context) error
	UpdateItem(ctx echo.Context) error
	HideItem(ctx echo.Context) error

	RegisterUser(ctx echo.Context) error

	HomePage(ctx echo.Context) error
}

type handler struct {
	userRepository  repository.UserRepository
	flatRepository  repository.FlatRepository
	groupRepository repository.GroupRepository
	itemRepository  repository.ItemRepository
	authService     service.AuthService
}

func NewHandler(user repository.UserRepository, flat repository.FlatRepository, group repository.GroupRepository, item repository.ItemRepository, auth service.AuthService) Handler {
	return &handler{
		userRepository:  user,
		flatRepository:  flat,
		groupRepository: group,
		itemRepository:  item,
		authService:     auth,
	}
}

func (h handler) HomePage(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
