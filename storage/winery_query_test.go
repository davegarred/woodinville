package storage

import (
	"testing"
	"github.com/davegarred/woodinville/domain"
	"github.com/stretchr/testify/assert"
)

var (
	wineryEL = &WineryQueryEventListener{}
	wineryName = "Winery Name"
)

func TestWineryQueryEventListener_OnWineryCreated(t *testing.T) {
	locations = make(map[domain.WineryId]*WineryQuery)

	wineryEL.OnWineryCreated(domain.WineryCreated{wineryId, wineryName})

	assert.Equal(t, wineryQuery(), locations[wineryId])
}

func TestWineryQueryEventListener_OnVisitAdded(t *testing.T) {
	locations = make(map[domain.WineryId]*WineryQuery)
	wineryEL.OnWineryCreated(domain.WineryCreated{wineryId, wineryName})

	wineryEL.OnVisitAdded(domain.VisitAdded{userId, aug06, wineryId})

	expected := wineryQuery()
	expected.Visits[userId] = []string{aug06}
	assert.Equal(t, expected, locations[wineryId])
}

func TestWineryQueryEventListener_OnWineryPositionUpdated(t *testing.T) {
locations = make(map[domain.WineryId]*WineryQuery)
	wineryEL.OnWineryCreated(domain.WineryCreated{wineryId, wineryName})

	wineryEL.OnWineryPositionUpdated(domain.WineryPositionUpdated{wineryId, 33.06, -112.13})

	expected := wineryQuery()
	expected.Lat = 33.06
	expected.Long = -112.13
	assert.Equal(t, expected, locations[wineryId])
}

func wineryQuery() *WineryQuery {
	return &WineryQuery{
		WineryId: wineryId,
		Name: wineryName,
		Visits: make(map[domain.UserId][]string),
	}
}