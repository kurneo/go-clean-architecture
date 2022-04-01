package v1

import (
	"kurneo/internal/auth/usercaes"
	"kurneo/internal/infrastructure/hash"
	"kurneo/internal/infrastructure/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(context echo.Context) error {
	username := context.FormValue("username")
	password := context.FormValue("password")

	if errorsValidate := validator.ValidateStruct(struct {
		Username string `validate:"required,alphanum"`
		Password string `validate:"required"`
	}{
		Username: username,
		Password: password,
	}); len(errorsValidate) > 0 {
		return context.JSON(http.StatusUnprocessableEntity,
			map[string]interface{}{
				"message": "The given data was invalid",
				"errors":  errorsValidate,
			},
		)
	}

	if token, user, errCreateToken := usercaes.NewAuthUsecase().Login(username, password); errCreateToken == nil {
		return context.JSON(http.StatusOK, map[string]interface{}{
			"token":   token,
			"message": "Login success",
			"user": map[string]interface{}{
				"id":    user.ID,
				"email": user.Email,
			},
		})
	} else if errCreateToken.Error() == "record not found" || errCreateToken.Error() == "password is not match" {
		return context.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Usename or password not match!",
		})
	} else {
		return context.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Server error!",
		})
	}
}

func SignUp(context echo.Context) error {
	username := context.FormValue("username")
	password := context.FormValue("password")
	email := context.FormValue("email")
	name := context.FormValue("name")
	gender := context.FormValue("gender")

	if errorsValidate := validator.ValidateStruct(struct {
		Username string `validate:"required,alphanum"`
		Password string `validate:"required"`
		Email    string `validate:"required,email"`
		Name     string `validate:"required"`
		Gender   string `validate:"required,oneof=male female"`
	}{
		Username: username,
		Password: password,
		Email:    email,
		Name:     name,
		Gender:   gender,
	}); len(errorsValidate) > 0 {
		return context.JSON(http.StatusUnprocessableEntity,
			map[string]interface{}{
				"message": "The given data was invalid",
				"errors":  errorsValidate,
			},
		)
	}

	authUsecase := usercaes.NewAuthUsecase()

	if check, errCheck := authUsecase.CheckUserExist(username); check == true {
		return context.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": "The usename is already used",
			},
		)
	} else if errCheck != nil && errCheck.Error() != "record not found" {
		return context.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Server error!",
		})
	}

	password, _ = hash.Make(password)

	if _, errSignUp := authUsecase.SignUp(map[string]interface{}{
		"Username": username,
		"Password": password,
		"Email":    email,
		"Name":     name,
		"Gender":   gender,
	}); errSignUp == nil {
		return context.JSON(http.StatusOK, map[string]interface{}{
			"message": "Signup success!",
		})
	} else {
		return context.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Server error!",
		})
	}
}
