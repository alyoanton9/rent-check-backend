package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"net/http"
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/model"
)

func (h handler) GetItems(ctx echo.Context) error {
	userId := ctx.Get("userId").(string)

	items, err := h.itemRepository.GetItems(userId)
	if err != nil {
		return HandleDbError(ctx, err, "error getting items")
	}

	itemResponses := lo.Map(items, func(item model.Item, _ int) dto.ItemResponse {
		return model.ItemToDto(item)
	})

	return ctx.JSON(http.StatusOK, itemResponses)
}
