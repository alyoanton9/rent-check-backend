package handler

import "github.com/labstack/echo/v4"

func (h handler) HideGroup(ctx echo.Context) error {
	userId := ctx.Get("userId").(string)

	groupId, err := ParsePathParamUInt64(ctx, "id")
	if err != nil {
		return err
	}

	err = h.groupRepository.HideGroup(groupId, userId)
	if err != nil {
		return HandleDbError(ctx, err, "error deleting group")
	}

	return nil
}
