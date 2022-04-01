package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testEventStruct struct {
	Event
}

func (event *testEventStruct) TriggerAll(argns ...interface{}) {
	for _, listener := range event.GetListeners() {
		listener.Handle(argns...)
	}
}

type testListenerStruct struct {
	ID string
}

func (listener *testListenerStruct) GetID() string {
	return listener.ID
}

func (listener *testListenerStruct) Handle(argns ...interface{}) {
	dispatched = true
}

var (
	testEvent  EventContract
	dispatched bool = false
)

func setUp() {
	event := Event{}
	event.SetName("test-event")
	testEvent = &testEventStruct{
		Event: event,
	}
}

func tearDown() {
	testEvent = nil
}

func TestRegisterListener(t *testing.T) {
	setUp()
	defer tearDown()

	testEvent.Register(&testListenerStruct{
		ID: "test-listener",
	})

	listeners := testEvent.GetListeners()

	assert.Equal(t, 1, len(listeners))
	assert.Equal(t, "test-listener", listeners[0].GetID())
}

func TestRemoveListener(t *testing.T) {
	setUp()
	defer tearDown()

	testEvent.Register(&testListenerStruct{
		ID: "test-listener",
	})

	testEvent.Deregister("test-listener")

	listeners := testEvent.GetListeners()

	assert.Equal(t, 0, len(listeners))
}

func TestDispatchEvent(t *testing.T) {
	setUp()
	defer tearDown()

	testEvent.Register(&testListenerStruct{
		ID: "test-listener",
	})
	testEvent.TriggerAll()
	assert.True(t, dispatched)
}
