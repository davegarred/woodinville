package storage

import (
	"github.com/davegarred/woodinville/domain"
)

type AreaQuery struct {
	domain.AreaId `json:"id"`
	Wineries            []domain.WineryId
	RecommendedWineries []RecommendedWinery
}

type RecommendedWinery struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	Zip     string `json:"zip"`
}

type AreaQueryEventListener struct{}

func (*AreaQueryEventListener) OnWineryCreated(e domain.WineryCreated) {
	q := FindArea()
	q.Wineries = append(q.Wineries, e.WineryId)
	//UpdateArea(q.AreaId, q)
}

func (*AreaQueryEventListener) OnAreaWineryRecommended(e domain.AreaWineryRecommended) {
	q := FindArea()
	wineries := q.RecommendedWineries
	if wineries == nil {
		wineries = make([]RecommendedWinery, 0)
	}
	q.RecommendedWineries = append(wineries, RecommendedWinery{
		e.Name,
		e.Address,
		e.City,
		e.Zip,
	})
}
