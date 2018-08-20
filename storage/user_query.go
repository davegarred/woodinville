package storage

import "github.com/davegarred/woodinville/domain"

type UserQuery struct {
	domain.UserId `json:"userId"`
	Name   string `json:"name"`
	Admin  bool `json:"admin"`
	Visits []WineryVisitQuery `json:"visits"`
}

type WineryVisitQuery struct {
	domain.WineryId
	Time string
}

func (q *UserQuery) OnUserCreated(e domain.UserCreated) {
	q.UserId = e.UserId
	q.Visits = make([]WineryVisitQuery,0)
}

func (q *UserQuery) OnUserIsSetAsAdmin(e domain.UserIsSetAsAdmin) {
	q.Admin = true
}

func (q *UserQuery) OnUserIsSetAsNotAdmin(e domain.UserIsSetAsNotAdmin) {
	q.Admin = false
}

func (q *UserQuery) OnVisitAdded(e domain.VisitAdded) {
	q.Visits = append(q.Visits, WineryVisitQuery{e.WineryId, e.Time})
}

