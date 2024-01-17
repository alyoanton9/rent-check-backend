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

	group.GET("/items", h.GetItems)
	group.POST("/items", h.CreateItem)
	group.PUT("/items/:id", h.UpdateItem)
	group.DELETE("/items/:id", h.HideItem)
	group.POST("/groups/:groupId/items", h.AddItemToGroup)
	group.DELETE("/group/:groupId/items", h.DeleteItemFromGroup)
	group.GET("/flats/:flatId/items", h.GetFlatItems)
	group.POST("/items/status", h.UpdateItemStatus)

	group.POST("/auth/register", h.Register)
	group.POST("/auth/login", h.Login)
	group.POST("/auth/logout", h.Logout)

	group.GET("", h.HomePage)
}
