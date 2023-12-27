package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"net/http"
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/model"
)

func (h handler) GetGroups(ctx echo.Context) error {
	userId := ctx.Get("userId").(string)

	idsString := ctx.QueryParam("ids")
	ids := ParseQueryParamList(idsString)

	groups, err := h.groupRepository.GetGroups(userId, ids)
	if err != nil {
		return HandleDbError(ctx, err, "error getting groups")
	}

	groupsResponse := lo.Map(groups, func(group model.Group, _ int) dto.GroupResponse {
		return model.GroupToDto(group)
	})

	return ctx.JSON(http.StatusOK, groupsResponse)
}
