package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/model"
)

func (h handler) UpdateItemStatus(ctx echo.Context) error {
	userId := ctx.Get("userId").(string)

	var itemStatusRequest *dto.ItemStatusRequest
	err := ParseBody(ctx, &itemStatusRequest, "item status request")
	if err != nil {
		return err
	}

	itemStatus, err := model.ParseStatus(itemStatusRequest.Status)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.itemRepository.UpdateItemStatus(
		itemStatusRequest.FlatId, itemStatusRequest.GroupId, itemStatusRequest.ItemId, itemStatus, userId)
	if err != nil {
		return HandleDbError(ctx, err, "error updating item status")
	}

	return nil
}
