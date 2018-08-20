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

type UserWineryQueryEventListener struct {}

func (*UserWineryQueryEventListener) OnUserCreated(e domain.UserCreated) {
	fmt.Println(&cqrs.QueryEventListener{})
	//q := FindUser(e.UserId)
	q := &UserQuery{}
	q.UserId = e.UserId
	q.Visits = make(map[domain.WineryId][]string)
	UpdateUser(e.UserId, q)
}

func (*UserWineryQueryEventListener) OnUserIsSetAsAdmin(e domain.UserIsSetAsAdmin) {
	q := FindUser(e.UserId)
	q.Admin = true
	UpdateUser(e.UserId, q)
}

func (*UserWineryQueryEventListener) OnUserIsSetAsNotAdmin(e domain.UserIsSetAsNotAdmin) {
	q := FindUser(e.UserId)
	q.Admin = false
	UpdateUser(e.UserId, q)
}

func (*UserWineryQueryEventListener) OnVisitAdded(e domain.VisitAdded) {
	q := FindUser(e.UserId)
	locationVisits := q.Visits[e.WineryId]
	if locationVisits == nil {
		locationVisits = make([]string,0)
	}
	locationVisits = append(locationVisits, e.Time)
	q.Visits[e.WineryId] = locationVisits
	UpdateUser(e.UserId, q)
}

//func (q *UserQuery) OnUserCreated(e domain.UserCreated) {
//	q.UserId = e.UserId
//	q.Visits = make(map[domain.WineryId][]string)
//}
//
//func (q *UserQuery) OnUserIsSetAsAdmin(e domain.UserIsSetAsAdmin) {
//	q.Admin = true
//}
//
//func (q *UserQuery) OnUserIsSetAsNotAdmin(e domain.UserIsSetAsNotAdmin) {
//	q.Admin = false
//}
//
//func (q *UserQuery) OnVisitAdded(e domain.VisitAdded) {
//	locationVisits := q.Visits[e.WineryId]
//	if locationVisits == nil {
//		locationVisits = make([]string,0)
//	}
//	locationVisits = append(locationVisits, e.Time)
//	q.Visits[e.WineryId] = locationVisits
//}
//
