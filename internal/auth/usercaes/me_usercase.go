package usercaes

import (
	"kurneo/internal/auth/infrastructure"
	authRepository "kurneo/internal/auth/repositories"
	"sync"

	"github.com/golang-jwt/jwt/v4"
)

type MeUsecase struct {
	repository *authRepository.User
}

var (
	meUsecase     infrastructure.MeUsecaseContract
	meUsecaseOnce sync.Once
)

func (usecase *MeUsecase) RefreshToken(token *jwt.Token) (string, error) {
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"]
	user, err := usecase.repository.FindByID(int(id.(float64)))

	if err != nil {
		return "", err
	}

	if token, errCreateToken := infrastructure.CreateToken(*user); errCreateToken == nil {
		return token, nil
	} else {
		return "", errCreateToken
	}
}

func (usecase *MeUsecase) GetProfile(token *jwt.Token) (*map[string]interface{}, error) {
	claims := token.Claims.(jwt.MapClaims)

	id := claims["id"]
	user, err := usecase.repository.FindByID(int(id.(float64)))

	if err != nil {
		return nil, err
	}

	var avatar interface{}

	if user.Avatar == "" {
		avatar = nil
	} else {
		avatar = user.Avatar
	}

	return &map[string]interface{}{
		"id":            user.ID,
		"dob":           user.DOB.Format("02-01-2006"),
		"about":         user.About,
		"avatar":        avatar,
		"name":          user.Name,
		"email":         user.Email,
		"gender":        user.Gender,
		"last_login_at": user.LastLoginAt.Format("02-01-2006 15:04"),
	}, nil
}

func NewMeUsecase() infrastructure.MeUsecaseContract {
	meUsecaseOnce.Do(func() {
		if meUsecase == nil {
			meUsecase = &MeUsecase{
				repository: authRepository.NewUserRepository(),
			}
		}
	})
	return meUsecase
}
