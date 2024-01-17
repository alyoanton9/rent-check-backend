package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/model"
)

func (h handler) CreateFlat(ctx echo.Context) error {
	userId := ctx.Get("userId").(uint64)

	var flatRequest *dto.FlatRequest
	err := ParseBody(ctx, &flatRequest, "flat request")
	if err != nil {
		return err
	}

	flat := model.DtoToFlat(*flatRequest, userId)

	err = h.flatRepository.CreateFlat(&flat)
	if err != nil {
		return HandleDbError(ctx, err, "error creating flat")
	}

	flatResponse := model.FlatToDto(flat)

	return ctx.JSON(http.StatusOK, flatResponse)
}
