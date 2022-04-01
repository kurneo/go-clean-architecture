package infrastructure

import (
	"kurneo/internal/auth/models"

	"github.com/golang-jwt/jwt/v4"
)

type AuthUsecaseContract interface {
	Login(username, password string) (string, *models.User, error)
	CheckUserExist(username string) (bool, error)
	SignUp(data map[string]interface{}) (*models.User, error)
}

type MeUsecaseContract interface {
	GetProfile(token *jwt.Token) (*map[string]interface{}, error)
	RefreshToken(token *jwt.Token) (string, error)
}
