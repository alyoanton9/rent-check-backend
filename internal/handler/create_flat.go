package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rent-checklist/internal/dto"
)

func (h handler) CreateFlat(ctx echo.Context) error {
	var flatRequest dto.FlatRequest
	if err := ctx.Bind(&flatRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	flat := flatRequest.ToModel()

	err := h.flatRepository.CreateFlat(&flat)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var flatResponse dto.FlatResponse
	(&flatResponse).FromModel(flat)

	return ctx.JSON(http.StatusOK, flatResponse)
}
