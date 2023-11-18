package route

import (
	"github.com/labstack/echo/v4"
	"rent-checklist-backend/internal/handler"
)

func Setup(e *echo.Echo, h handler.Handler) {
	group := e.Group("/api/v1")

	group.GET("/flats", h.GetFlats)
	group.POST("/flats", h.CreateFlat)

	group.GET("", h.HomePage)
}
