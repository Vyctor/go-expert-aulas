package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e TestEvent) GetName() string {
	return e.Name
}

func (e TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h TestEventHandler) Handle(event EventInterface) {

}

type EventDispatcherTestSuit struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuit) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}
	suite.event = TestEvent{
		Name:    "test",
		Payload: "test",
	}
	suite.event2 = TestEvent{
		Name:    "test2",
		Payload: "test2",
	}
}

func (suite *EventDispatcherTestSuit) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), suite.handler3)
	suite.Nil(err)
	suite.Equal(3, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1])
	assert.Equal(suite.T(), suite.handler3, suite.eventDispatcher.handlers[suite.event.GetName()][2])
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuit))
}
