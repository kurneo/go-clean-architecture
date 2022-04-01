package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field         string      `json:"field"`
	Rule          string      `json:"rule"`
	ProvidedValue interface{} `json:"provided_value"`
}

func ValidateStruct(stc interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	validate.RegisterTagNameFunc(func(structField reflect.StructField) string {
		name := strings.SplitN(structField.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	err := validate.Struct(stc)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = strings.ToLower(err.Field())
			element.Rule = err.Tag()
			element.ProvidedValue = err.Value()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateValue(name string, value interface{}, rules string) *ErrorResponse {
	validate := validator.New()

	err := validate.Var(value, rules)

	if err != nil {
		var errorValidate = &ErrorResponse{}
		err := err.(validator.ValidationErrors)[0]
		errorValidate.Field = name
		errorValidate.Rule = err.Tag()
		errorValidate.ProvidedValue = value
		return errorValidate
	}

	return nil
}
