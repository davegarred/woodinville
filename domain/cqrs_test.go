package domain

import (
	"testing"
	"github.com/davegarred/woodinville/cqrs"
	"fmt"
)

var (
	createSome = &CreateSomeCommand{"an id"}
	nameSome   = &NameSomeCommand{"a name"}
	genericSome = &GenericUpdateSomeCommand{}
)

func TestCommandHandler(t *testing.T) {
	someAggregate := &SomeAggregate{}
	commandGateway := cqrs.NewCommandGateway(someAggregate)
	commandGateway.Register(someAggregate)

	dispatch(commandGateway, createSome)
	dispatch(commandGateway, nameSome)
	//commandGateway.Dispatch(genericSome)

	fmt.Printf("Final aggregate state: %+v\n", *someAggregate)
}

func dispatch(commandGateway *cqrs.CommandGateway, c interface{}) error {
	err := commandGateway.Dispatch(c)
	if err != nil {
		panic(err)
	}
	return nil
}
