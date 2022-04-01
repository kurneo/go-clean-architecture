package event

import (
	"kurneo/internal/auth/models"
	"kurneo/internal/auth/repositories"
	"kurneo/internal/infrastructure/event"
	"sync"
	"time"

	"github.com/labstack/gommon/log"
)

// Structs
type AuthLoginEvent struct {
	event.Event
}

type AuthLoginListener struct {
	ID string
}

// Variables
var (
	authLoginEvent     *AuthLoginEvent
	authLoginEventOnce sync.Once
)

// Methods
func (event *AuthLoginEvent) TriggerAll(argns ...interface{}) {
	for _, listener := range event.GetListeners() {
		listener.Handle(argns...)
	}
}

func (listener *AuthLoginListener) GetID() string {
	return listener.ID
}

func (listener *AuthLoginListener) Handle(argns ...interface{}) {
	user := argns[0].(*models.User)
	_, err := repositories.NewUserRepository().Update(user, map[string]interface{}{"LastLoginAt": time.Now()})
	if err != nil {
		log.Error(err)
	}
}

func NewAuthLoginEvent() event.EventContract {
	authLoginEventOnce.Do(func() {
		if authLoginEvent == nil {
			event := event.Event{}
			event.SetName("auth.login-event")
			authLoginEvent = &AuthLoginEvent{
				Event: event,
			}
		}
	})
	return authLoginEvent
}
