package query

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
