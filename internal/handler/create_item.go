package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/model"
)

func (h handler) CreateItem(ctx echo.Context) error {
	userId := ctx.Get("userId").(string)

	var itemRequest *dto.ItemRequest
	err := ParseBody(ctx, &itemRequest, "item request")
	if err != nil {
		return err
	}

	item := model.DtoToItem(*itemRequest, userId)

	err = h.itemRepository.CreateItem(&item)
	if err != nil {
		return HandleDbError(ctx, err, "error creating item")
	}

	itemResponse := model.ItemToDto(item)

	return ctx.JSON(http.StatusOK, itemResponse)
}
