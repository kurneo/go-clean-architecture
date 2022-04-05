package infrastructure

import (
	"kurneo/internal/infrastructure/paginator"
	"kurneo/internal/infrastructure/repositories"
	"kurneo/internal/tag/models"
)

type TagUsecaseContract interface {
	
	GetList(
		conditions []repositories.Condition,
		with []string,
		selectColumns []string,
		orders []repositories.Order,
		paginate *repositories.Paginate,
	) (*[]models.Tag, *paginator.Paginator, error)
	
	Get(id int) (*models.Tag, error)
	
	Create(data map[string]interface{}) (*models.Tag, error)
	
	Update(tag *models.Tag, data map[string]interface{}) (bool, error)
	
	Delete(tag *models.Tag) (bool, error)
}
