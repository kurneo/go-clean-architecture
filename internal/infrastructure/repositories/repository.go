package repositories

import (
	"kurneo/internal/infrastructure/paginator"
	"reflect"

	"gorm.io/gorm"
)

type Repository[T any] struct {
	Model T
	DB    *gorm.DB
}

func (repository Repository[T]) NewModel() T {
	newModel := reflect.New(reflect.TypeOf(repository.Model)).Elem().Interface().(T)
	return newModel
}

func (repository *Repository[T]) All(with []string, selectColumns []string, orders []Order, paginate *Paginate) (*[]T, *paginator.Paginator, error) {
	var list []T

	query := repository.DB.Model(repository.NewModel())

	applySelectColumns(query, selectColumns)
	applyEagerLoad(query, with)
	applyOrder(query, orders)
	applyPaginate(paginate, query)

	var totalCount int64
	if errorCount := query.Count(&totalCount).Error; errorCount != nil {
		return nil, nil, errorCount
	}

	if errorQuery := query.Find(&list).Error; errorQuery != nil {
		return nil, nil, errorQuery
	}

	if pg, _ := paginator.Populate(paginate.Page, paginate.Limit, int(totalCount)); pg != nil {
		return &list, pg, nil
	}

	return &list, nil, nil
}

func (repository *Repository[T]) AllBy(conditions []Condition, with []string, selectColumns []string, orders []Order, paginate *Paginate) (*[]T, *paginator.Paginator, error) {
	var list []T

	query := repository.DB.Model(repository.NewModel())

	var totalCount int64
	if errorCount := query.Count(&totalCount).Error; errorCount != nil {
		return nil, nil, errorCount
	}

	applySelectColumns(query, selectColumns)
	applyEagerLoad(query, with)
	applyOrder(query, orders)
	applyCondition(query, conditions)
	applyPaginate(paginate, query)

	if errorQuery := query.Find(&list).Error; errorQuery != nil {
		return nil, nil, errorQuery
	}

	if pg, _ := paginator.Populate(paginate.Page, paginate.Limit, int(totalCount)); pg != nil {
		return &list, pg, nil
	}

	return &list, nil, nil
}

func (repository *Repository[T]) FirstBy(conditions []Condition, with []string, selectColumns []string) (*T, error) {
	model := repository.NewModel()
	query := repository.DB.Model(model)
	applyCondition(query, conditions)
	applyEagerLoad(query, with)
	applySelectColumns(query, selectColumns)

	if errorGet := query.First(&model).Error; errorGet == nil {
		return &model, nil
	} else {
		return nil, errorGet
	}
}

func (repository *Repository[T]) FindByID(id int) (*T, error) {
	var model = repository.NewModel()
	if errorGet := repository.DB.Where("id = ?", id).First(&model).Error; errorGet == nil {
		return &model, nil
	} else {
		return nil, errorGet
	}
}

func (repository *Repository[T]) Store(data map[string]interface{}) (*T, error) {
	model := repository.NewModel()
	modelValue := reflect.ValueOf(&model)
	for key, value := range data {
		switch modelValue.Elem().FieldByName(key).Kind() {
		case reflect.String:
			modelValue.Elem().FieldByName(key).SetString(value.(string))
		case reflect.Int:
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Int64:
			modelValue.Elem().FieldByName(key).SetInt(value.(int64))
		case reflect.Bool:
			modelValue.Elem().FieldByName(key).SetBool(value.(bool))
		default:
		}
	}
	if e := repository.DB.Create(&model).Error; e != nil {
		return nil, e
	}
	return &model, nil
}

func (repository *Repository[T]) Update(model *T, data map[string]interface{}) (bool, error) {
	if e := repository.DB.Model(model).Updates(removeEmptyFields(data)).Error; e != nil {
		return false, e
	}
	return true, nil
}

func (repository *Repository[T]) Delete(model *T) (bool, error) {
	if e := repository.DB.Delete(model).Error; e != nil {
		return false, e
	}
	return true, nil
}

func (repository *Repository[T]) Count(conditions []Condition) (int64, error) {
	var count int64
	query := repository.DB.Model(repository.NewModel())
	applyCondition(query, conditions)
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
