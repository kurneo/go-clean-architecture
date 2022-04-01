package repositories

import (
	"kurneo/internal/infrastructure/paginator"
)

type RepositoryContract[T any] interface {
	NewModel() T
	All(with []string, selectColumns []string, orders []Order, paginate *Paginate) (*[]T, *paginator.Paginator, error)
	AllBy(conditions []Condition, with []string, selectColumns []string, orders []Order, paginate *Paginate) (*[]T, *paginator.Paginator, error)
	FirstBy(conditions []Condition, with []string, selectColumns []string) (*T, error)
	FindByID(id int) (*T, error)
	Store(data map[string]interface{}) (*T, error)
	Update(tag *T, data map[string]interface{}) (bool, error)
	Delete(tag *T) (bool, error)
	Count(conditions []Condition) (int64, error)
}
