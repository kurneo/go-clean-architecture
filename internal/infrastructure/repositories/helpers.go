package repositories

import (
	"kurneo/internal/infrastructure/paginator"

	"gorm.io/gorm"
)

type Condition struct {
	Column   string
	Operator string
	Value    interface{}
}

type Order struct {
	Column string
	Dir    string
}

type Paginate struct {
	Page  int
	Limit int
}

func applyCondition(query *gorm.DB, conditions []Condition) {
	if len(conditions) == 0 {
		return
	}
	for _, condition := range conditions {
		if condition.Column == "" && condition.Operator == "" {
			continue
		}
		query = query.Where(condition.Column+" "+condition.Operator+" ?", condition.Value)
	}
}

func applyOrder(query *gorm.DB, orders []Order) {
	if len(orders) == 0 {
		return
	}
	for _, order := range orders {
		if order.Dir == "" && order.Column == "" {
			continue
		}
		query = query.Order(order.Column + " " + order.Dir)
	}
}

func applyEagerLoad(query *gorm.DB, with []string) {
	if len(with) == 0 {
		return
	}
	for _, relation := range with {
		query = query.Preload(relation)
	}
}

func applySelectColumns(query *gorm.DB, columns []string) {
	if len(columns) > 0 {
		query = query.Select(columns)
	}
}

func applyPaginate(paginate *Paginate, query *gorm.DB) {
	if paginate != nil && paginate.Page != 0 && paginate.Limit != 0 {
		if offset, err := paginator.ResolveOffset(paginate.Page, paginate.Limit); err == nil {
			query.Offset(offset).Limit(paginate.Limit)
		}
	}
}

func removeEmptyFields(inputs map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{}

	for key, value := range inputs {
		if value != "" {
			data[key] = value
		}
	}

	return data
}
