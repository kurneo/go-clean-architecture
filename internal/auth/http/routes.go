package http

import (
	controllersV1 "kurneo/internal/auth/http/controllers/v1"
	"kurneo/internal/infrastructure/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(group *echo.Group) {
	v1Group := group.Group("/admin/v1")
	v1Group.POST("/auth/login", controllersV1.Login)
	v1Group.POST("/auth/signup", controllersV1.SignUp)

	v1Group.GET("/auth/me", controllersV1.Me, middlewares.JwtMiddleware())
	v1Group.POST("/auth/refresh-token", controllersV1.RefreshToken, middlewares.JwtMiddleware())
}
