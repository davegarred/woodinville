package cqrs

import (
	"testing"
)

type CreateUser struct {

}
func (c *CreateUser) Handle() []*EventWrapper {
	return nil
}


func TestCommandHandler(t *testing.T) {
	testAggregate := NewAggregate()
	createUser := CreateUser{}
	testAggregate.Handle(createUser)
}
