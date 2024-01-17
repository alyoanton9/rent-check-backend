package handler

import (
	"github.com/labstack/echo/v4"
	"rent-checklist-backend/internal/dto"
)

func (h handler) AddGroupToFlat(ctx echo.Context) error {
	userId := ctx.Get("userId").(uint64)

	flatId, err := ParsePathParamUInt64(ctx, "flatId")
	if err != nil {
		return err
	}

	var groupIdRequest *dto.GroupIdRequest
	err = ParseBody(ctx, &groupIdRequest, "group id request")
	if err != nil {
		return err
	}

	err = h.groupRepository.AddGroupToFlat(groupIdRequest.Id, flatId, userId)
	if err != nil {
		return HandleDbError(ctx, err, "error adding group to flat")
	}

	return nil
}
