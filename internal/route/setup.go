package route

import (
	"github.com/labstack/echo/v4"
	"rent-checklist/internal/handler"
)

func Setup(e *echo.Echo, h handler.Handler) {
	group := e.Group("/api/v1")

	group.GET("", h.HomePage)
}