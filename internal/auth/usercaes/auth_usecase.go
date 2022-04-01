package usercaes

import (
	"errors"
	"kurneo/internal/auth/event"
	"kurneo/internal/auth/infrastructure"
	"kurneo/internal/auth/models"
	authRepository "kurneo/internal/auth/repositories"
	"kurneo/internal/infrastructure/repositories"
	"sync"
)

type AuthUsecase struct {
	repository *authRepository.User
}

var (
	authUsecase     infrastructure.AuthUsecaseContract
	authUsecaseOnce sync.Once
)

func (usercase *AuthUsecase) Login(username, password string) (string, *models.User, error) {
	user, err := usercase.repository.FirstBy(
		[]repositories.Condition{
			{
				Column:   "username",
				Operator: "=",
				Value:    username,
			},
		},
		make([]string, 0),
		[]string{"id", "username", "email", "password"},
	)

	if err != nil {
		return "", nil, err
	}

	if user.VerifyPassword(password) == false {
		return "", nil, errors.New("password is not match")
	}

	if token, errCreateToken := infrastructure.CreateToken(*user); errCreateToken == nil {
		event.NewAuthLoginEvent().TriggerAll(user)
		return token, user, nil
	} else {
		return "", nil, errCreateToken
	}
}

func (usercase *AuthUsecase) SignUp(data map[string]interface{}) (*models.User, error) {
	user, err := usercase.repository.Store(data)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (usercase *AuthUsecase) CheckUserExist(username string) (bool, error) {
	_, err := usercase.repository.FirstBy(
		[]repositories.Condition{
			{
				Column:   "username",
				Operator: "=",
				Value:    username,
			},
		},
		make([]string, 0),
		[]string{"id"},
	)
	if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

func NewAuthUsecase() infrastructure.AuthUsecaseContract {
	authUsecaseOnce.Do(func() {
		if authUsecase == nil {
			authUsecase = &AuthUsecase{
				repository: authRepository.NewUserRepository(),
			}
		}
	})
	return authUsecase
}
