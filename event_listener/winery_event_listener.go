package event_listener

import (
	"github.com/davegarred/woodinville/domain"
	"github.com/davegarred/woodinville/query"
	"github.com/davegarred/woodinville/storage"
)

type WineryQueryEventListener struct {}

func (*WineryQueryEventListener) OnWineryCreated(e domain.WineryCreated) {
	q := &query.WineryQuery{}
	q.WineryId = e.WineryId
	q.Name = e.Name
	q.Visits = make(map[domain.UserId][]string)
	storage.UpdateLocation(e.WineryId, q)
}

func (*WineryQueryEventListener) OnWineryPositionUpdated(e domain.WineryPositionUpdated) {
	q := storage.FindLocation(e.WineryId)
	q.Lat = e.Lat
	q.Long = e.Long
	storage.UpdateLocation(e.WineryId, q)
}

func (*WineryQueryEventListener) OnWineryAddressUpdated(e domain.WineryAddressUpdated) {
	q := storage.FindLocation(e.WineryId)
	q.Address = e.Address
	q.City = e.City
	q.Zip = e.Zip
	storage.UpdateLocation(e.WineryId, q)
}

func (*WineryQueryEventListener) OnVisitAdded(e domain.VisitAdded) {
	q := storage.FindLocation(e.WineryId)
	if q == nil {
		return
	}
	locationVisits := q.Visits[e.UserId]
	if locationVisits == nil {
		locationVisits = make([]string,0)
	}
	locationVisits = append(locationVisits, e.Time)
	q.Visits[e.UserId] = locationVisits
	storage.UpdateLocation(e.WineryId, q)
}