package repositories

import (
	"kurneo/internal/infrastructure/dbconn"
	"kurneo/internal/infrastructure/repositories"
	"kurneo/internal/tag/models"
	"sync"
)

type Tag struct {
	repositories.Repository[models.Tag]
}

var (
	tagRepoOnce sync.Once
	tagRepo     *Tag
)

func NewTagRepository() *Tag {
	tagRepoOnce.Do(func() {
		if tagRepo == nil {
			conn, _ := dbconn.NewConnection()
			tagRepo = &Tag{
				repositories.Repository[models.Tag]{
					Model: models.Tag{},
					DB:    conn.DB,
				},
			}
		}
	})
	return tagRepo
}
