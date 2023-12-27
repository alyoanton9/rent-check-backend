package handler

import (
	"github.com/labstack/echo/v4"
	"rent-checklist-backend/internal/dto"
)

func (h handler) DeleteItemFromGroup(ctx echo.Context) error {
	userId := ctx.Get("userId").(string)

	groupId, err := ParsePathParamUInt64(ctx, "groupId")
	if err != nil {
		return err
	}

	var flatItemIdsRequest *dto.FlatItemIdsRequest
	err = ParseBody(ctx, &flatItemIdsRequest, "flat item ids request")
	if err != nil {
		return err
	}

	err = h.itemRepository.DeleteItemFromGroup(flatItemIdsRequest.FlatId, groupId, flatItemIdsRequest.ItemId, userId)
	if err != nil {
		return HandleDbError(ctx, err, "error deleting item from flat group")
	}

	return nil
}
