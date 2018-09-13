package event_listener

import (
	"github.com/davegarred/woodinville/domain"
	"github.com/davegarred/woodinville/query"
	"github.com/davegarred/woodinville/storage"
)


type UserQueryEventListener struct {}

func (*UserQueryEventListener) OnUserCreated(e domain.UserCreated) {
	q := &query.UserQuery{}
	q.UserId = e.UserId
	q.Name = e.Name
	q.Visits = make(map[domain.WineryId][]string)
	storage.UpdateUser(e.UserId, q)
}

func (*UserQueryEventListener) OnUserIsSetAsAdmin(e domain.UserIsSetAsAdmin) {
	q := storage.FindUser(e.UserId)
	q.Admin = true
	storage.UpdateUser(e.UserId, q)
}

func (*UserQueryEventListener) OnUserIsSetAsNotAdmin(e domain.UserIsSetAsNotAdmin) {
	q := storage.FindUser(e.UserId)
	q.Admin = false
	storage.UpdateUser(e.UserId, q)
}

func (*UserQueryEventListener) OnVisitAdded(e domain.VisitAdded) {
	q := storage.FindUser(e.UserId)
	if q == nil {
		return
	}
	if q.Visits == nil {
		q.Visits = make(map[domain.WineryId][]string)
	}
	locationVisits := q.Visits[e.WineryId]
	if locationVisits == nil {
		locationVisits = make([]string,0)
	}
	locationVisits = append(locationVisits, e.Time)
	q.Visits[e.WineryId] = locationVisits
	storage.UpdateUser(e.UserId, q)
}
