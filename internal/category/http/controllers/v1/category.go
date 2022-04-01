package controllers

import (
	"kurneo/internal/category/usecases"
	"kurneo/internal/infrastructure/controllers"
	"kurneo/internal/infrastructure/repositories"
	"kurneo/internal/infrastructure/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ListCategories(context echo.Context) error {
	status := context.QueryParam("status")
	search := context.QueryParam("search")

	page, limit, errValidatePaginate := controllers.GetPaginateParams(context)
	sortField, sortDir, errValidateSort := controllers.GetSortParams(context)

	errValidate := validator.ValidateStruct(struct {
		Status string `validate:"omitempty,oneof=public draft"`
	}{Status: status})

	errValidate = append(errValidate, errValidatePaginate...)
	errValidate = append(errValidate, errValidateSort...)

	if len(errValidate) > 0 {
		return context.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": "The given data was invalid",
			"errors":  errValidate,
		})
	}

	usecase := usecases.NewCategoryUsecase()

	conditions := []repositories.Condition{}
	if search != "" {
		conditions = append(
			conditions,
			repositories.Condition{
				Column:   "name",
				Operator: "ilike",
				Value:    "%" + search + "%",
			},
		)
	}

	if status != "" {
		conditions = append(
			conditions,
			repositories.Condition{
				Column:   "status",
				Operator: "=",
				Value:    status,
			},
		)
	}

	if categories, paginate, err := usecase.GetList(
		conditions,
		make([]string, 0),
		make([]string, 0),
		[]repositories.Order{{Dir: sortDir, Column: sortField}},
		&repositories.Paginate{Page: page, Limit: limit},
	); err == nil {
		context.Response().Header().Set("X-Total-Count", strconv.Itoa(paginate.Total))
		context.Response().Header().Set("X-Total-Page", strconv.Itoa(paginate.TotalPages))
		return context.JSON(http.StatusOK, categories)
	} else {
		return context.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"message": "Ops, have an error"},
		)
	}
}

func GetCategory(context echo.Context) error {
	id := context.Param("id")

	errValidate := validator.ValidateStruct(struct {
		ID string `validate:"numeric,gte=1"`
	}{ID: id})

	if len(errValidate) > 0 {
		return context.JSON(
			http.StatusUnprocessableEntity,
			map[string]interface{}{"message": "The given data was invalid", "errors": errValidate},
		)
	}

	intID, _ := strconv.Atoi(id)

	if model, err := usecases.NewCategoryUsecase().Get(intID); err != nil {
		if err.Error() == "record not found" {
			return context.JSON(
				http.StatusBadRequest,
				map[string]interface{}{"message": "Category is invalid"},
			)
		} else {
			return context.JSON(
				http.StatusInternalServerError,
				map[string]interface{}{"message": "Ops, have an error!"},
			)
		}
	} else {
		return context.JSON(http.StatusUnprocessableEntity, model)
	}
}

func StoreCategory(context echo.Context) error {
	body := struct {
		Name        string `validate:"required" json:"name"`
		Description string `json:"description"`
		IsDefault   bool   `json:"is_default"`
		Status      string `validate:"required,oneof=public draft" json:"status"`
	}{}

	if errBindBody := context.Bind(&body); errBindBody != nil {
		return context.JSON(
			http.StatusBadRequest,
			map[string]interface{}{"message": "Bad request"},
		)
	}

	if errorsValidate := validator.ValidateStruct(body); len(errorsValidate) > 0 {
		return context.JSON(
			http.StatusUnprocessableEntity,
			map[string]interface{}{
				"message": "The given data was invalid",
				"errors":  errorsValidate,
			},
		)
	}

	if category, err := usecases.NewCategoryUsecase().Create(map[string]interface{}{
		"Name":        body.Name,
		"Description": body.Description,
		"IsDefault":   body.IsDefault,
		"Status":      body.Status,
	}); err == nil {
		return context.JSON(http.StatusOK, category)
	} else {
		return context.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"message": "Create category failed"},
		)
	}
}

func UpdateCategory(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return context.JSON(
			http.StatusBadRequest,
			map[string]interface{}{"message": "Bad request"},
		)
	}

	data := map[string]interface{}{}
	var errorsValidate []interface{}

	if name := context.FormValue("name"); name != "" {
		data["name"] = name
	}

	if description := context.FormValue("description"); description != "" {
		data["description"] = description
	}

	if status := context.FormValue("status"); status != "" {
		errorValidateStatus := validator.ValidateValue(
			"status",
			context.FormValue("status"),
			"oneof=public draft",
		)
		if errorValidateStatus == nil {
			data["status"] = status
		} else {
			errorsValidate = append(errorsValidate, errorValidateStatus)
		}
	}

	if len(errorsValidate) > 0 {
		return context.JSON(
			http.StatusUnprocessableEntity,
			map[string]interface{}{
				"message": "The given data was invalid",
				"errors":  errorsValidate,
			},
		)
	}

	usecase := usecases.NewCategoryUsecase()
	model, err := usecase.Get(id)
	if err != nil {
		return context.JSON(
			http.StatusBadRequest,
			map[string]interface{}{"message": "Category is invalid"},
		)
	}

	if success, _ := usecase.Update(model, data); success {
		return context.JSON(http.StatusOK, model)
	} else {
		return context.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"message": "Update category failed"},
		)
	}
}

func DeleteCategory(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return context.JSON(
			http.StatusBadRequest,
			map[string]interface{}{"message": "Bad request"},
		)
	}

	usecase := usecases.NewCategoryUsecase()

	model, err := usecase.Get(id)

	if err != nil {
		if err.Error() == "record not found" {
			return context.JSON(
				http.StatusBadRequest,
				map[string]interface{}{"message": "Category is invalid"},
			)
		} else {
			return context.JSON(
				http.StatusInternalServerError,
				map[string]interface{}{"message": "Ops, have an error!"},
			)
		}
	}

	if success, _ := usecase.Delete(model); success {
		return context.JSON(http.StatusOK, true)
	} else {
		return context.JSON(http.StatusInternalServerError, false)
	}
}
