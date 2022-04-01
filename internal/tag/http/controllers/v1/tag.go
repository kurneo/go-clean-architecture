package controllers

import (
	"fmt"
	"kurneo/internal/infrastructure/controllers"
	"kurneo/internal/infrastructure/repositories"
	"kurneo/internal/infrastructure/validator"
	"kurneo/internal/tag/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// List godoc
// @Summary      list tags
// @Description  Get list tags
// @Tags         Tags
// @Accept       json
// @Produce      json
// @Param   	 page     query    int     false    "Page number"         		  default(1)
// @Param   	 limit    query    int     fase     "Number of items per page" 	  default(10)
// @Success      200  {string} ok
// @Failure      400  {string}  string
// @Router       /api/admin/v1/tags/list [get]
// @Security Bearer
func ListTags(context echo.Context) error {

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

	if tags, paginator, err := usecases.NewTagUsecase().GetList(
		conditions,
		make([]string, 0),
		make([]string, 0),
		[]repositories.Order{{Dir: sortDir, Column: sortField}},
		&repositories.Paginate{Page: page, Limit: limit},
	); err == nil {
		context.Response().Header().Set("X-Total-Count", strconv.Itoa(paginator.Total))
		context.Response().Header().Set("X-Total-Pages", strconv.Itoa(paginator.TotalPages))
		return context.JSON(http.StatusOK, tags)
	} else {
		return context.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"message": "Ops, have an error!"},
		)
	}
}

func GetTag(context echo.Context) error {
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

	if tag, err := usecases.NewTagUsecase().Get(intID); err != nil {
		if err.Error() == "record not found" {
			return context.JSON(
				http.StatusBadRequest,
				map[string]interface{}{"message": "Tag is invalid"},
			)
		} else {
			return context.JSON(
				http.StatusInternalServerError,
				map[string]interface{}{"message": "Ops, have an error!"},
			)
		}
	} else {
		return context.JSON(http.StatusOK, tag)
	}
}

func StoreTag(context echo.Context) error {

	body := struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
		Status      string `json:"status" validate:"required,oneof=public draft"`
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

	if tag, err := usecases.NewTagUsecase().Create(map[string]interface{}{
		"Name":        body.Name,
		"Description": body.Description,
		"Status":      body.Status,
	}); err == nil {
		return context.JSON(http.StatusOK, tag)
	} else {
		return context.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{
				"message": "Create tag failed",
			},
		)
	}
}

func UpdateTag(context echo.Context) error {
	id := context.Param("id")
	name := context.FormValue("name")
	description := context.FormValue("description")
	status := context.FormValue("status")

	errValidate := validator.ValidateStruct(struct {
		ID          string `validate:"numeric,gte=1"`
		Name        string `validate:"omitempty"`
		Description string `validate:"omitempty"`
		Status      string `validate:"omitempty,oneof=public draft"`
	}{ID: id, Name: name, Description: description, Status: status})

	if len(errValidate) > 0 {
		return context.JSON(
			http.StatusUnprocessableEntity,
			map[string]interface{}{"message": "The given data was invalid", "errors": errValidate},
		)
	}

	fmt.Println(status)

	intID, _ := strconv.Atoi(id)

	usecase := usecases.NewTagUsecase()
	model, errGet := usecase.Get(intID)

	if errGet != nil {
		if errGet.Error() == "record not found" {
			return context.JSON(
				http.StatusBadRequest,
				map[string]interface{}{"message": "Tag is invalid"},
			)
		} else {
			return context.JSON(
				http.StatusInternalServerError,
				map[string]interface{}{"message": "Ops, have an error!"},
			)
		}
	}

	if success, _ := usecase.Update(model, map[string]interface{}{
		"Name":        name,
		"Description": description,
		"Status":      status,
	}); success {
		return context.JSON(http.StatusOK, model)
	} else {
		return context.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"message": "Updated tag failed"},
		)
	}
}

func DeleteTag(context echo.Context) error {
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
	usecase := usecases.NewTagUsecase()

	model, err := usecase.Get(intID)

	if err != nil {
		if err.Error() == "record not found" {
			return context.JSON(
				http.StatusBadRequest,
				map[string]interface{}{"message": "Tag is invalid"},
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
