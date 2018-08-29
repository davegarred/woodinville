package event_listener

import (
	"testing"
	"github.com/davegarred/woodinville/domain"
	"github.com/stretchr/testify/assert"
	"github.com/davegarred/woodinville/query"
	"github.com/davegarred/woodinville/storage"
)

var (
	wineryEL = &WineryQueryEventListener{}
	wineryName = "Winery Name"
)

func TestWineryQueryEventListener_OnWineryCreated(t *testing.T) {
	storage.ResetTests()

	wineryEL.OnWineryCreated(domain.WineryCreated{wineryId, wineryName})

	assert.Equal(t, wineryQuery(), storage.FindLocation(wineryId))
}

func TestWineryQueryEventListener_OnVisitAdded(t *testing.T) {
	storage.ResetTests()
	wineryEL.OnWineryCreated(domain.WineryCreated{wineryId, wineryName})

	wineryEL.OnVisitAdded(domain.VisitAdded{userId, aug06, wineryId})

	expected := wineryQuery()
	expected.Visits[userId] = []string{aug06}
	assert.Equal(t, expected, storage.FindLocation(wineryId))
}

func TestWineryQueryEventListener_OnWineryPositionUpdated(t *testing.T) {
	storage.ResetTests()
	wineryEL.OnWineryCreated(domain.WineryCreated{wineryId, wineryName})

	wineryEL.OnWineryPositionUpdated(domain.WineryPositionUpdated{wineryId, 33.06, -112.13})

	expected := wineryQuery()
	expected.Lat = 33.06
	expected.Long = -112.13
	assert.Equal(t, expected, storage.FindLocation(wineryId))
}

func wineryQuery() *query.WineryQuery {
	return &query.WineryQuery{
		WineryId: wineryId,
		Name: wineryName,
		Visits: make(map[domain.UserId][]string),
	}
}