package server

import (
	"kurneo/internal/auth"
	"kurneo/internal/category"
	"kurneo/internal/tag"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(app *echo.Echo) {
	app.GET("/swagger/doc.json", func(c echo.Context) error {
		return c.File("docs/swagger.json")
	})
	app.GET("/swagger/*", echoSwagger.WrapHandler)

	group := app.Group("/api")
	auth.RegisterModule(app, group)
	tag.RegisterModule(app, group)
	category.RegisterModule(app, group)
}
