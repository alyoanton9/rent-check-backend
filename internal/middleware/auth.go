package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"rent-checklist-backend/internal/repository"
	"rent-checklist-backend/internal/service"
)

func MakeAuthConfig(authService service.AuthService, userRepository repository.UserRepository) middleware.KeyAuthConfig {
	var authConfig middleware.KeyAuthConfig

	authConfig.AuthScheme = authService.GetAuthScheme()

	authConfig.Validator = func(authToken string, ctx echo.Context) (bool, error) {
		user, err := userRepository.GetUserByAuthToken(authToken)
		if err != nil {
			log.Printf("error getting user by auth token: %s", err.Error())

			return false, err
		}

		ctx.Set("userId", user.Id)

		return true, nil
	}

	authConfig.Skipper = func(ctx echo.Context) bool {
		return ctx.Request().URL.Path == "/api/v1/register"
	}

	return authConfig
}
