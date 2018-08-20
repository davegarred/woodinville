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
	area := FindArea()
	assert.Equal(t, 1, len(area))
}

func TestSerialization(t *testing.T) {
	area := FindArea()
	ser,err := json.Marshal(area)
	assert.Nil(t, err)
	assert.Equal(t, `[{"id":"DAR","lat":47.73222,"long":-122.14273,"name":"Darby","address":"14450 Redmond-Woodinville Rd NE","city":"Woodinville","zip":"98072"}]`, string(ser))
}