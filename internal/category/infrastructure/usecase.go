package infrastructure

import (
	"kurneo/internal/category/models"
	"kurneo/internal/infrastructure/paginator"
	"kurneo/internal/infrastructure/repositories"
)

type CategoryUsecaseContract interface {
	GetList(
		conditions []repositories.Condition,
		with []string,
		selectColumns []string,
		orders []repositories.Order,
		paginate *repositories.Paginate,
	) (*[]models.Category, *paginator.Paginator, error)
	Get(id int) (*models.Category, error)
	Create(data map[string]interface{}) (*models.Category, error)
	Update(tag *models.Category, data map[string]interface{}) (bool, error)
	Delete(tag *models.Category) (bool, error)
}
