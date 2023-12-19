package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/model"
)

func (h handler) CreateGroup(ctx echo.Context) error {
	userId := ctx.Get("userId").(string)

	var groupRequest *dto.GroupRequest
	err := ParseBody(ctx, &groupRequest, "group request")
	if err != nil {
		return err
	}

	group := model.DtoToGroup(*groupRequest, userId)

	err = h.groupRepository.CreateGroup(&group)
	if err != nil {
		return HandleDbError(ctx, err, "error creating group")
	}

	groupResponse := model.GroupToDto(group)

	return ctx.JSON(http.StatusOK, groupResponse)
}
