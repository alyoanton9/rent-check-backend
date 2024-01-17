package handler

import "github.com/labstack/echo/v4"

func (h handler) Logout(ctx echo.Context) error {
	userId := ctx.Get("userId").(uint64)

	user, err := h.userRepository.GetUserById(userId)
	if err != nil {
		return HandleDbError(ctx, err, "error when logout user")
	}

	user.AuthToken = nil
	err = h.userRepository.UpdateUser(user)
	if err != nil {
		return HandleDbError(ctx, err, "error when deleting user's auth token")
	}

	return nil
}
