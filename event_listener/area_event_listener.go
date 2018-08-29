package event_listener

import (
	"github.com/davegarred/woodinville/domain"
	"github.com/davegarred/woodinville/query"
	"github.com/davegarred/woodinville/storage"
)

type AreaQueryEventListener struct{}

func (*AreaQueryEventListener) OnWineryCreated(e domain.WineryCreated) {
	q := storage.FindArea()
	q.Wineries = append(q.Wineries, e.WineryId)
	//UpdateArea(q.AreaId, q)
}

func (*AreaQueryEventListener) OnAreaWineryRecommended(e domain.AreaWineryRecommended) {
	q := storage.FindArea()
	wineries := q.RecommendedWineries
	if wineries == nil {
		wineries = make([]query.RecommendedWinery, 0)
	}
	q.RecommendedWineries = append(wineries, query.RecommendedWinery{
		e.Name,
		e.Address,
		e.City,
		e.Zip,
	})
}
