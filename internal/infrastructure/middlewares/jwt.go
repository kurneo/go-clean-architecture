package middlewares

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"strconv"
	"time"
)

func JwtMiddleware() echo.MiddlewareFunc {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	config := middleware.JWTConfig{
		TokenLookup: "header:" + echo.HeaderAuthorization,
		ParseTokenFunc: func(auth string, context echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return signingKey, nil
			}

			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}

			if !token.Valid {
				return nil, errors.New("invalid token")
			}

			claims := token.Claims.(jwt.MapClaims)

			ext := fmt.Sprintf("%v", claims["expired_at"])

			extParseInt, errParse := strconv.ParseInt(ext, 10, 64)

			if errParse != nil {
				return nil, errors.New("parse expire time failed")
			}

			if extParseInt-time.Now().Unix() < 60 {
				return nil, errors.New("token is expired")
			}

			return token, nil
		},
	}
	return middleware.JWTWithConfig(config)
}
