package handler

import (
	"github.com/labstack/echo/v4"
)

func (h handler) DeleteFlat(ctx echo.Context) error {
	userId := ctx.Get("userId").(string)

	flatId, err := ParsePathParamUInt64(ctx, "id")
	if err != nil {
		return err
	}

	err = h.flatRepository.DeleteFlat(flatId, userId)
	if err != nil {
		return HandleDbError(ctx, err, "error deleting flat")
	}

	return nil
}
