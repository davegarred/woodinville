package query

import (
	"github.com/davegarred/woodinville/domain"
)

type WineryQuery struct {
	domain.WineryId `json:"id"`
	Lat     float32 `json:"lat"`
	Long    float32 `json:"long"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	City    string  `json:"city"`
	Zip     string  `json:"zip"`
	Visits map[domain.UserId][]string `json:"visits"`
}
