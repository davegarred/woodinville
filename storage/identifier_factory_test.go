package storage

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIdentifierFactory_Next(t *testing.T) {
	factory := NewIdentifierFactory(IdentifierType("WL"))
	assert.Equal(t, "WL10001", factory.Next())
	assert.Equal(t, "WL10002", factory.Next())
	assert.Equal(t, "WL10003", factory.Next())
}
