package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/model"
)

func (h handler) UpdateFlat(ctx echo.Context) error {
	userId := ctx.Get("userId").(string)

	flatId, err := ParsePathParamUInt64(ctx, "id")
	if err != nil {
		return err
	}

	var flatRequest *dto.FlatRequest
	err = ParseBody(ctx, &flatRequest, "flat request")
	if err != nil {
		return err
	}

	flat := model.DtoToFlat(*flatRequest, "")
	flat.Id = flatId

	err = h.flatRepository.UpdateFlat(&flat, userId)

	if err != nil {
		return HandleDbError(ctx, err, "error updating flat")
	}

	flatResponse := model.FlatToDto(flat)

	return ctx.JSON(http.StatusOK, flatResponse)
}
