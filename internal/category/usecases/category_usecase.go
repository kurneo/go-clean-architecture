package usecases

import (
	"kurneo/internal/category/infrastructure"
	"kurneo/internal/category/models"
	"kurneo/internal/category/repositories"
	"kurneo/internal/infrastructure/paginator"
	infraRepository "kurneo/internal/infrastructure/repositories"
	"sync"
)

type CategoryUsecase struct {
	repository repositories.Category
}

var (
	categoryUsecase     infrastructure.CategoryUsecaseContract
	categoryUsecaseOnce sync.Once
)

func (usecase *CategoryUsecase) GetList(
	conditions []infraRepository.Condition,
	with []string,
	selectColumns []string,
	orders []infraRepository.Order,
	paginate *infraRepository.Paginate,
) (*[]models.Category, *paginator.Paginator, error) {
	return usecase.repository.AllBy(conditions, with, selectColumns, orders, paginate)
}

func (usecase *CategoryUsecase) Get(id int) (*models.Category, error) {
	return usecase.repository.FindByID(id)
}
func (usecase *CategoryUsecase) Create(data map[string]interface{}) (*models.Category, error) {
	return usecase.repository.Store(data)
}
func (usecase *CategoryUsecase) Update(tag *models.Category, data map[string]interface{}) (bool, error) {
	return usecase.repository.Update(tag, data)
}
func (usecase *CategoryUsecase) Delete(tag *models.Category) (bool, error) {
	return usecase.repository.Delete(tag)
}

func NewCategoryUsecase() infrastructure.CategoryUsecaseContract {
	categoryUsecaseOnce.Do(func() {
		if categoryUsecase == nil {
			categoryUsecase = &CategoryUsecase{
				repository: *repositories.NewCategoryRepository(),
			}
		}
	})

	return categoryUsecase
}
