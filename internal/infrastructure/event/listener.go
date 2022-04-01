package event

type ListenerContract interface {
	Handle(data ...interface{})
	GetID() string
}
