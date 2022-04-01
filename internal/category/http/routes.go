package http

import (
	"kurneo/internal/category/http/controllers/v1"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(group *echo.Group) {
	v1Group := group.Group("/admin/v1")
	tagGroup := v1Group.Group("/categories")
	tagGroup.GET("/list", controllers.ListCategories)
	tagGroup.GET("/:id", controllers.GetCategory)
	tagGroup.POST("/store", controllers.StoreCategory)
	tagGroup.PUT("/update/:id", controllers.UpdateCategory)
	tagGroup.DELETE("/delete/:id", controllers.DeleteCategory)
}
