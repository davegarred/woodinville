package event_listener

import (
	"github.com/davegarred/cqrs"
	"github.com/davegarred/cqrs/components"
	"github.com/davegarred/cqrs/persist"
	"github.com/davegarred/woodinville/domain"
)

var eventStore cqrs.EventStore
var commandGateway *components.CommandGateway

func init() {

	eventBus := components.NewEventBus()
	eventStore = persist.NewMemEventStore(eventBus)
	commandGateway = components.NewCommandGateway(eventStore)
	commandGateway.RegisterAggregate(&domain.User{})
	commandGateway.RegisterAggregate(&domain.Winery{})
	commandGateway.RegisterAggregate(&domain.Area{})
	eventBus.RegisterQueryEventHandlers(&UserQueryEventListener{})
	eventBus.RegisterQueryEventHandlers(&WineryQueryEventListener{})
	eventBus.RegisterQueryEventHandlers(&AreaQueryEventListener{})
}

func Dispatch(command cqrs.Command) {
	commandGateway.Dispatch(command)
}
