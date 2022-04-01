package auth

import (
	"kurneo/internal/auth/event"
	"kurneo/internal/auth/http"

	"github.com/labstack/echo/v4"
)

func RegisterModule(app *echo.Echo, group *echo.Group) {
	http.RegisterRoutes(group)
	registerEvent()
}

func registerEvent() {
	loginEvent := event.NewAuthLoginEvent()
	loginEvent.Register(&event.AuthLoginListener{
		ID: "auth.login-listener",
	})
}
