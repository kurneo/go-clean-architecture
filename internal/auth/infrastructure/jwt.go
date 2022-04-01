package infrastructure

import (
	"errors"
	"kurneo/internal/auth/models"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(user models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		return "", errors.New("jwt secret mismatch")
	}

	claims := jwt.MapClaims{
		"id":         user.ID,
		"username":   user.Username,
		"name":       user.Name,
		"expired_at": createExpiredAt(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func createExpiredAt() string {
	return strconv.FormatInt(time.Now().Add(time.Hour*24).Unix(), 10)
}
