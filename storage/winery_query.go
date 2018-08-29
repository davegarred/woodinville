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
	Visits map[domain.UserId][]string `json:"visits"`
}

type WineryQueryEventListener struct {}

func (*WineryQueryEventListener) OnWineryCreated(e domain.WineryCreated) {
	q := &WineryQuery{}
	q.WineryId = e.WineryId
	q.Name = e.Name
	q.Visits = make(map[domain.UserId][]string)
	UpdateLocation(e.WineryId, q)
}

func (*WineryQueryEventListener) OnWineryPositionUpdated(e domain.WineryPositionUpdated) {
	q := FindLocation(e.WineryId)
	q.Lat = e.Lat
	q.Long = e.Long
	UpdateLocation(e.WineryId, q)
}

func (*WineryQueryEventListener) OnWineryAddressUpdated(e domain.WineryAddressUpdated) {
	q := FindLocation(e.WineryId)
	q.Address = e.Address
	q.City = e.City
	q.Zip = e.Zip
	UpdateLocation(e.WineryId, q)
}

func (*WineryQueryEventListener) OnVisitAdded(e domain.VisitAdded) {
	q := FindLocation(e.WineryId)
	if q == nil {
		return
	}
	locationVisits := q.Visits[e.UserId]
	if locationVisits == nil {
		locationVisits = make([]string,0)
	}
	locationVisits = append(locationVisits, e.Time)
	q.Visits[e.UserId] = locationVisits
	UpdateLocation(e.WineryId, q)
}