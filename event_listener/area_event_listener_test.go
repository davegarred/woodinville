package event_listener

import (
	"testing"
	"github.com/davegarred/woodinville/domain"
	"github.com/stretchr/testify/assert"
	"github.com/davegarred/woodinville/storage"
)

var (
	wineryId = domain.WineryId("a_winery_id")
	areaEL = &AreaQueryEventListener{}
)

func TestAreaQueryEventListener_OnWineryCreated(t *testing.T) {
	storage.ResetTests()
	assert.Equal(t, 0, len(storage.FindArea().Wineries))

	areaEL.OnWineryCreated(domain.WineryCreated{wineryId, wineryName})

	resultArea := storage.FindArea()
	assert.Equal(t, 1, len(resultArea.Wineries))
	assert.Equal(t, wineryId, resultArea.Wineries[0])
}

func TestAreaQueryEventListener_OnAreaWineryRecommended(t *testing.T) {
	storage.ResetTests()

	areaEL.OnAreaWineryRecommended(domain.AreaWineryRecommended{domain.AreaId("SEA"), wineryName, "11234 Road St.", "Woodinville", "98111" })

	assert.Equal(t, 1, len(storage.FindArea().RecommendedWineries))
}
