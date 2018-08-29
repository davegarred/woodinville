package storage

import (
	"testing"
	"github.com/davegarred/woodinville/domain"
	"github.com/stretchr/testify/assert"
)

var (
	wineryId = domain.WineryId("a_winery_id")
	areaEL = &AreaQueryEventListener{}
)

func TestAreaQueryEventListener_OnWineryCreated(t *testing.T) {
	areaQuery.Wineries = []domain.WineryId{}
	assert.Equal(t, 0, len(FindArea().Wineries))

	areaEL.OnWineryCreated(domain.WineryCreated{wineryId, wineryName})

	resultArea := FindArea()
	assert.Equal(t, 1, len(resultArea.Wineries))
	assert.Equal(t, wineryId, resultArea.Wineries[0])
}

func TestAreaQueryEventListener_OnAreaWineryRecommended(t *testing.T) {
	areaQuery.Wineries = []domain.WineryId{}

	areaEL.OnAreaWineryRecommended(domain.AreaWineryRecommended{domain.AreaId("SEA"), wineryName, "11234 Road St.", "Woodinville", "98111" })

	assert.Equal(t, 1, len(areaQuery.RecommendedWineries))
}
