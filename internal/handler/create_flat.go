package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"rent-checklist-backend/internal/dto"
	e "rent-checklist-backend/internal/error"
)

func (h handler) CreateFlat(ctx echo.Context) error {
	var flatRequest dto.FlatRequest
	if err := ctx.Bind(&flatRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	flat := dto.ToModel(flatRequest)

	err := h.flatRepository.CreateFlat(&flat)
	if err != nil {
		log.Printf("error creating flat: %v", err.Error())

		var keyAlreadyExist *e.KeyAlreadyExist
		if errors.As(err, &keyAlreadyExist) {
			errorResponse := e.UniqueErrorResponse{Field: keyAlreadyExist.Field, Msg: keyAlreadyExist.Msg}
			return ctx.JSON(http.StatusBadRequest, errorResponse)
		}

		return echo.ErrInternalServerError
	}

	flatResponse := dto.FromModel(flat)

	return ctx.JSON(http.StatusOK, flatResponse)
}
