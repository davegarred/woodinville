package storage

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
}


func (*UserWineryQueryEventListener) OnWineryCreated(e domain.WineryCreated) {
	q := &WineryQuery{}
	q.WineryId = e.WineryId
	q.Name = e.Name
	UpdateLocation(e.WineryId, q)
}

func (*UserWineryQueryEventListener) OnWineryPositionUpdated(e domain.WineryPositionUpdated) {
	q := FindLocation(e.WineryId)
	q.Lat = e.Lat
	q.Long = e.Long
	UpdateLocation(e.WineryId, q)
}

func (*UserWineryQueryEventListener) OnWineryAddressUpdated(e domain.WineryAddressUpdated) {
	q := FindLocation(e.WineryId)
	q.Address = e.Address
	q.City = e.City
	q.Zip = e.Zip
	UpdateLocation(e.WineryId, q)
}
