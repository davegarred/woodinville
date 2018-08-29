package event_listener

import (
	"github.com/davegarred/cqrs"
	"github.com/davegarred/woodinville/domain"
)

var eventStore cqrs.EventStore
var commandGateway *cqrs.CommandGateway

func init() {

	eventStore = cqrs.NewMemEventStore()
	commandGateway = cqrs.NewCommandGateway(eventStore)
	commandGateway.RegisterAggregate(&domain.User{})
	commandGateway.RegisterAggregate(&domain.Winery{})
	commandGateway.RegisterAggregate(&domain.Area{})
	commandGateway.RegisterQueryEventHandlers(&UserQueryEventListener{})
	commandGateway.RegisterQueryEventHandlers(&WineryQueryEventListener{})
	commandGateway.RegisterQueryEventHandlers(&AreaQueryEventListener{})
}

func Dispatch(command cqrs.Command) {
	commandGateway.Dispatch(command)
}
