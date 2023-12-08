package handler

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	e "rent-checklist-backend/internal/error"
	"strconv"
)

func HandleDbError(ctx echo.Context, err error, logPrefix string) error {
	log.Printf("%s: %s", logPrefix, err.Error())

	switch err.(type) {
	case *e.KeyNotFound:
		return ctx.JSON(http.StatusBadRequest, err)
	case *e.KeyAlreadyExist:
		return ctx.JSON(http.StatusBadRequest, err)
	case *e.ForbiddenAction:
		return ctx.JSON(http.StatusForbidden, err)
	default:
		return echo.ErrInternalServerError
	}
}

func ParsePathParamUInt64(ctx echo.Context, paramName string) (uint64, error) {
	idStr := ctx.Param(paramName)
	paramValue, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Printf("error parsing %s: %s", paramName, err.Error())

		return 0, echo.ErrBadRequest
	}

	return paramValue, nil
}

func ParseBody[T any](ctx echo.Context, dto *T, dtoName string) error {
	if err := ctx.Bind(&dto); err != nil {
		log.Printf("error parsing %s body: %s", dtoName, err.Error())

		return echo.ErrBadRequest
	}

	return nil
}