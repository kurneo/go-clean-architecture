package usecases

import (
	"kurneo/internal/infrastructure/paginator"
	infraRepository "kurneo/internal/infrastructure/repositories"
	"kurneo/internal/tag/infrastructure"
	"kurneo/internal/tag/models"
	"kurneo/internal/tag/repositories"
	"sync"
)

type TagUsecase struct {
	repository repositories.Tag
}

var (
	tagUsecase     infrastructure.TagUsecaseContract
	tagUsecaseOnce sync.Once
)

func (usecase *TagUsecase) GetList(
	conditions []infraRepository.Condition,
	with []string,
	selectColumns []string,
	orders []infraRepository.Order,
	paginate *infraRepository.Paginate,
) (*[]models.Tag, *paginator.Paginator, error) {
	return usecase.repository.AllBy(conditions, with, selectColumns, orders, paginate)
}

func (usecase *TagUsecase) Get(id int) (*models.Tag, error) {
	return usecase.repository.FindByID(id)
}
func (usecase *TagUsecase) Create(data map[string]interface{}) (*models.Tag, error) {
	return usecase.repository.Store(data)
}
func (usecase *TagUsecase) Update(tag *models.Tag, data map[string]interface{}) (bool, error) {
	return usecase.repository.Update(tag, data)
}
func (usecase *TagUsecase) Delete(tag *models.Tag) (bool, error) {
	return usecase.repository.Delete(tag)
}

func NewTagUsecase() infrastructure.TagUsecaseContract {
	tagUsecaseOnce.Do(func() {
		if tagUsecase == nil {
			tagUsecase = &TagUsecase{
				repository: *repositories.NewTagRepository(),
			}
		}
	})

	return tagUsecase
}
