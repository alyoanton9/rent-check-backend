package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"log"
	"net/http"
	"rent-checklist/internal/dto"
	"rent-checklist/internal/model"
)

func (h handler) GetFlats(ctx echo.Context) error {
	flats, err := h.flatRepository.GetFlats()
	if err != nil {
		log.Printf("error getting flats: %v", err.Error())

		return echo.ErrInternalServerError
	}

	flatsResponse := lo.Map(flats, func(flat model.Flat, _ int) dto.FlatResponse {
		return dto.FromModel(flat)
	})

	return ctx.JSON(http.StatusOK, flatsResponse)
}
