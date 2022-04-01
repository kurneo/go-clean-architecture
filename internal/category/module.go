package category

import (
	"kurneo/internal/category/http"

	"github.com/labstack/echo/v4"
)

func RegisterModule(app *echo.Echo, group *echo.Group) {
	http.RegisterRoutes(group)
}
