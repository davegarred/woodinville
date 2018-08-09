package domain

import (
	"github.com/davegarred/cqrs"
	"errors"
)

type SomeAggregate struct {
	someId string
	name string
}

func (a *SomeAggregate) HandleCreateSome(e *CreateSomeCommand) ([]*cqrs.EventWrapper,error) {
	a.someId = e.Id
	return []*cqrs.EventWrapper{}, nil
}

func (a *SomeAggregate) HandleNameSome(e *NameSomeCommand) ([]*cqrs.EventWrapper,error) {
	if a.someId == "" {
		return nil, errors.New("aggregate has not been initialized")
	}
	a.name = e.Name
	return []*cqrs.EventWrapper{}, nil
}
