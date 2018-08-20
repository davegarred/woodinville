package storage

import "github.com/davegarred/woodinville/domain"

type UserQuery struct {
	domain.UserId `json:"id"`
	Name   string `json:"name"`
	Admin  bool `json:"admin"`
	Visits map[domain.WineryId][]string `json:"visits"`
}

func (q *UserQuery) OnUserCreated(e domain.UserCreated) {
	q.UserId = e.UserId
	q.Visits = make(map[domain.WineryId][]string)
}

func (q *UserQuery) OnUserIsSetAsAdmin(e domain.UserIsSetAsAdmin) {
	q.Admin = true
}

func (q *UserQuery) OnUserIsSetAsNotAdmin(e domain.UserIsSetAsNotAdmin) {
	q.Admin = false
}

func (q *UserQuery) OnVisitAdded(e domain.VisitAdded) {
	locationVisits := q.Visits[e.WineryId]
	if locationVisits == nil {
		locationVisits = make([]string,0)
	}
	locationVisits = append(locationVisits, e.Time)
	q.Visits[e.WineryId] = locationVisits
}

