package http

import (
	"kurneo/internal/infrastructure/middlewares"
	"kurneo/internal/tag/http/controllers/v1"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(group *echo.Group) {
	v1Group := group.Group("/admin/v1", middlewares.JwtMiddleware())
	tagGroup := v1Group.Group("/tags")
	tagGroup.GET("/list", controllers.ListTags)
	tagGroup.GET("/:id", controllers.GetTag)
	tagGroup.POST("/store", controllers.StoreTag)
	tagGroup.PUT("/update/:id", controllers.UpdateTag)
	tagGroup.DELETE("/delete/:id", controllers.DeleteTag)
}
