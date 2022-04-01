package v1

import (
	"kurneo/internal/auth/usercaes"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func Me(context echo.Context) error {
	token := context.Get("user").(*jwt.Token)
	if data, err := usercaes.NewMeUsecase().GetProfile(token); err == nil {
		return context.JSON(http.StatusOK, data)
	} else {
		return context.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Server error!",
		})
	}
}

func RefreshToken(context echo.Context) error {
	jwtToken := context.Get("user").(*jwt.Token)
	token, err := usercaes.NewMeUsecase().RefreshToken(jwtToken)
	if err == nil {
		return context.JSON(http.StatusOK, map[string]interface{}{
			"token":   token,
			"message": "Refresh token success",
		})
	} else {
		return context.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Server error!",
		})
	}
}
