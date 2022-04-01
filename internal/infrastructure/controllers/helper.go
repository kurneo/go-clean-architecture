package controllers

import (
	"kurneo/internal/infrastructure/validator"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetPaginateParams(context echo.Context) (int, int, []*validator.ErrorResponse) {
	page := context.QueryParam("page")
	limit := context.QueryParam("limit")

	errorsValidate := validator.ValidateStruct(struct {
		Page  string `validate:"omitempty,numeric,gte=1" json:"page"`
		Limit string `validate:"omitempty,numeric,gte=1" json:"limit"`
	}{
		Page:  page,
		Limit: limit,
	})

	if len(errorsValidate) > 0 {
		return 0, 0, errorsValidate
	}

	if page == "" {
		page = "1"
	}

	if limit == "" {
		limit = "10"
	}

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)

	return intPage, intLimit, errorsValidate
}

func GetSortParams(context echo.Context) (string, string, []*validator.ErrorResponse) {
	sortField := context.QueryParam("sort_field")
	sortDir := context.QueryParam("sort_dir")

	errorsValidate := validator.ValidateStruct(struct {
		SortField string `validate:"omitempty" json:"sort_field"`
		SortDir   string `validate:"omitempty,oneof=asc desc" json:"sort_dir"`
	}{
		SortField: sortField,
		SortDir:   sortDir,
	})

	if len(errorsValidate) > 0 {
		return "", "", errorsValidate
	}

	if sortField == "" {
		sortField = "id"
	}

	if sortDir == "" {
		sortDir = "asc"
	}

	return sortField, sortDir, errorsValidate
}
