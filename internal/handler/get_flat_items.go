package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rent-checklist-backend/internal/model"
)

func (h handler) GetFlatItems(ctx echo.Context) error {
	userId := ctx.Get("userId").(string)

	flatId, err := ParsePathParamUInt64(ctx, "flatId")
	if err != nil {
		return err
	}

	groupItemsList, err := h.itemRepository.GetFlatItems(flatId, userId)
	if err != nil {
		return HandleDbError(ctx, err, "error getting flat items")
	}

	return ctx.JSON(http.StatusOK, model.GroupItemsToDto(groupItemsList))
}
