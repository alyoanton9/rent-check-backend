package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"net/http"
	"rent-checklist-backend/internal/dto"
	"rent-checklist-backend/internal/model"
)

func (h handler) GetFlats(ctx echo.Context) error {
	userId := ctx.Get("userId").(string)

	flats, err := h.flatRepository.GetFlats(userId)
	if err != nil {
		return HandleDbError(ctx, err, "error getting flats")
	}

	flatsResponse := lo.Map(flats, func(flat model.Flat, _ int) dto.FlatResponse {
		return model.FlatToDto(flat)
	})

	return ctx.JSON(http.StatusOK, flatsResponse)
}
