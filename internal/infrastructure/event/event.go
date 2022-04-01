package event

import "fmt"

type EventContract interface {
	GetName() string
	SetName(name string)
	GetListeners() []ListenerContract
	Register(listener ListenerContract)
	Deregister(ID string)
	TriggerAll(argns ...interface{})
}

type Event struct {
	listeners []ListenerContract
	name      string
}

func removeFromslice(listeners []ListenerContract, ID string) []ListenerContract {
	length := len(listeners)
	for index, listener := range listeners {
		if listener.GetID() == ID {
			listeners[index] = listeners[length-1]
			return listeners[:length-1]
		}
	}
	return listeners
}

func (event *Event) GetName() string {
	return event.name
}

func (event *Event) SetName(name string) {
	event.name = name
}

func (event *Event) Register(listener ListenerContract) {
	event.listeners = append(event.listeners, listener)
}

func (event *Event) Deregister(ID string) {
	fmt.Println(removeFromslice(event.listeners, ID))
	event.listeners = removeFromslice(event.listeners, ID)
}

func (event *Event) GetListeners() []ListenerContract {
	return event.listeners
}
