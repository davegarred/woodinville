package cqrs

import (
	"time"
	"reflect"
)

type Command interface {
	Handle() []*EventWrapper
}

type EventWrapper struct {
	AggregateId string
	AggregateType string
	DateTime time.Time
	Payload interface{}
}

type Aggregate struct {
	commandHandlers map[reflect.Type]Command
}
func NewAggregate() *Aggregate {
	return &Aggregate{make(map[reflect.Type]Command)}
}
//func (a *Aggregate) Register(command Command) {
//	commandType := reflect.TypeOf(command)
//	a.commandHandlers[commandType] = command
//}

func (a *Aggregate) Handle(command Command) []EventWrapper {
	return command.Handle()
}