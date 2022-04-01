package repositories

import (
	"kurneo/internal/auth/models"
	"kurneo/internal/infrastructure/dbconn"
	"kurneo/internal/infrastructure/repositories"
	"sync"
)

type User struct {
	repositories.Repository[models.User]
}

var (
	userRepositoryOnce sync.Once
	userRepository     *User
)

func NewUserRepository() *User {
	userRepositoryOnce.Do(func() {
		if userRepository == nil {
			conn, _ := dbconn.NewConnection()
			userRepository = &User{
				Repository: repositories.Repository[models.User]{
					Model: models.User{},
					DB:    conn.DB,
				},
			}
		}
	})

	return userRepository
}
