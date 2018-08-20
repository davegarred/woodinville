package storage

import (
	"github.com/davegarred/woodinville/domain"
)

type WineryQuery struct {
	domain.WineryId `json:"wineryId"`
	Lat     float32 `json:"lat"`
	Long    float32 `json:"long"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	City    string  `json:"city"`
	Zip     string  `json:"zip"`
}


func (q *WineryQuery) OnWineryCreated(e domain.WineryCreated) {
	q.WineryId = e.WineryId
	q.Name = e.Name
}

func (q *WineryQuery) OnWineryPositionUpdated(e domain.WineryPositionUpdated) {
	q.Lat = e.Lat
	q.Long = e.Long
}

func (q *WineryQuery) OnWineryAddressUpdated(e domain.WineryAddressUpdated) {
	q.Address = e.Address
	q.City = e.City
	q.Zip = e.Zip
}
