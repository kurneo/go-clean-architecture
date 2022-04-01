package server

import (
	"kurneo/internal/infrastructure/middlewares"
	"os"

	"github.com/labstack/echo/v4"
)

func Start(app *echo.Echo) error {

	//Register middleware
	app.Use(middlewares.RateLimiterMiddleware())
	app.Use(middlewares.GzipMiddleware())
	RegisterRoutes(app)

	//Start application
	return app.Start(os.Getenv("HTTP_PORT"))
}
