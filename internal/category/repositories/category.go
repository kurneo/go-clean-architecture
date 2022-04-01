package repositories

import (
	"kurneo/internal/category/models"
	"kurneo/internal/infrastructure/dbconn"
	"kurneo/internal/infrastructure/repositories"
	"sync"
)

type Category struct {
	repositories.Repository[models.Category]
}

var (
	categoryRepoOnce sync.Once
	categoryRepo     *Category
)

func NewCategoryRepository() *Category {
	categoryRepoOnce.Do(func() {
		if categoryRepo == nil {
			conn, _ := dbconn.NewConnection()
			categoryRepo = &Category{
				repositories.Repository[models.Category]{
					Model: models.Category{},
					DB:    conn.DB,
				},
			}
		}
	})
	return categoryRepo
}
