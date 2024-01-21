package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"net/http"
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/model"
)

func (h handler) GetGroupFlats(ctx echo.Context) error {
	userId := ctx.Get("userId").(uint64)

	groupId, err := ParsePathParamUInt64(ctx, "id")
	if err != nil {
		return err
	}

	flats, err := h.flatRepository.GetFlatsByGroupId(groupId, userId)
	if err != nil {
		return HandleDbError(ctx, err, "error getting flats by groupId")
	}

	flatsResponse := lo.Map(flats, func(flat model.Flat, _ int) dto.FlatResponse {
		return model.FlatToDto(flat)
	})

	return ctx.JSON(http.StatusOK, flatsResponse)
}
