package route

import (
	"github.com/labstack/echo/v4"
	"rent-checklist-backend/internal/handler"
)

func Setup(e *echo.Echo, h handler.Handler) {
	group := e.Group("/api/v1")

	group.GET("/flats", h.GetFlats)
	group.POST("/flats", h.CreateFlat)
	group.PUT("/flats/:id", h.UpdateFlat)
	group.DELETE("/flats/:id", h.DeleteFlat)

	group.GET("/groups", h.GetGroups)
	group.POST("/groups", h.CreateGroup)
	group.PUT("/groups/:id", h.UpdateGroup)
	group.POST("/flats/:flatId/groups", h.AddGroupToFlat)
	group.DELETE("/flats/:flatId/groups", h.DeleteGroupFromFlat)
	group.DELETE("/groups/:id", h.HideGroup)

	group.POST("/register", h.RegisterUser)

	group.GET("", h.HomePage)

}
