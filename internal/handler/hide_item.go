package handler

import "github.com/labstack/echo/v4"

func (h handler) HideItem(ctx echo.Context) error {
	userId := ctx.Get("userId").(string)

	itemId, err := ParsePathParamUInt64(ctx, "id")
	if err != nil {
		return err
	}

	err = h.itemRepository.HideItem(itemId, userId)
	if err != nil {
		return HandleDbError(ctx, err, "error deleting item")
	}

	return nil
}
