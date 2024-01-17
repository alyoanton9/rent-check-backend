package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/model"
)

func (h handler) UpdateGroup(ctx echo.Context) error {
	userId := ctx.Get("userId").(uint64)

	groupId, err := ParsePathParamUInt64(ctx, "id")
	if err != nil {
		return err
	}

	var groupRequest *dto.GroupRequest
	err = ParseBody(ctx, &groupRequest, "group request")
	if err != nil {
		return err
	}

	group := model.DtoToGroup(*groupRequest, 0)
	group.Id = groupId

	err = h.groupRepository.UpdateGroup(&group, userId)

	if err != nil {
		return HandleDbError(ctx, err, "error updating group")
	}

	groupResponse := model.GroupToDto(group)

	return ctx.JSON(http.StatusOK, groupResponse)
}
