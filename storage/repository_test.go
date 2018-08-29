package storage

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestFindLocation(t *testing.T) {
	darby := FindLocation("DAR")
	assert.NotNil(t, darby)

	other := FindLocation("something else")
	assert.Nil(t, other)
}

func TestFindArea(t *testing.T) {
	area := FindWineries()
	assert.Equal(t, 1, len(area))
}

func TestSerialization(t *testing.T) {
	area := FindWineries()
	ser,err := json.Marshal(area)
	assert.Nil(t, err)
	assert.Equal(t, `[{"id":"DAR","lat":47.7318,"long":-122.14036,"name":"Darby","address":"14450 Redmond-Woodinville Rd NE","city":"Woodinville","zip":"98072","visits":null}]`, string(ser))
}

func TestNextWineryIdentifier(t *testing.T) {
	wineryIdentifierFactory = NewIdentifierFactory(wineryIdentifierType)
	assert.Equal(t, "WL10001", wineryIdentifierFactory.Next())
}