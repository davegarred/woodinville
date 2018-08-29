package storage

import (
	"github.com/davegarred/woodinville/domain"
	"github.com/davegarred/cqrs"
	"fmt"
)

type UserQuery struct {
	domain.UserId `json:"id"`
	Name   string `json:"name"`
	Admin  bool `json:"admin"`
	Visits map[domain.WineryId][]string `json:"visits"`
}

type UserQueryEventListener struct {}

func (*UserQueryEventListener) OnUserCreated(e domain.UserCreated) {
	fmt.Println(&cqrs.QueryEventListener{})
	q := &UserQuery{}
	q.UserId = e.UserId
	q.Name = e.Name
	q.Visits = make(map[domain.WineryId][]string)
	UpdateUser(e.UserId, q)
}

func (*UserQueryEventListener) OnUserIsSetAsAdmin(e domain.UserIsSetAsAdmin) {
	q := FindUser(e.UserId)
	q.Admin = true
	UpdateUser(e.UserId, q)
}

func (*UserQueryEventListener) OnUserIsSetAsNotAdmin(e domain.UserIsSetAsNotAdmin) {
	q := FindUser(e.UserId)
	q.Admin = false
	UpdateUser(e.UserId, q)
}

func (*UserQueryEventListener) OnVisitAdded(e domain.VisitAdded) {
	q := FindUser(e.UserId)
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
	UpdateUser(e.UserId, q)
}
