package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/model"
)

func (h handler) UpdateItem(ctx echo.Context) error {
	userId := ctx.Get("userId").(string)

	itemId, err := ParsePathParamUInt64(ctx, "id")
	if err != nil {
		return err
	}

	var itemRequest *dto.ItemRequest
	err = ParseBody(ctx, &itemRequest, "item request")
	if err != nil {
		return err
	}

	item := model.DtoToItem(*itemRequest, "")
	item.Id = itemId

	err = h.itemRepository.UpdateItem(&item, userId)

	if err != nil {
		return HandleDbError(ctx, err, "error updating item")
	}

	itemResponse := model.ItemToDto(item)

	return ctx.JSON(http.StatusOK, itemResponse)
}
