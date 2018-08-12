package domain

import "github.com/davegarred/cqrs"

type UserId string

type User struct {
	userId UserId
	name string
	visits []WineryVisit
}

type WineryVisit struct {
	WineryId
	time string
}

func (w *Winery) HandleCreateUser(c CreateUser) ([]cqrs.Event,error) {
	return []cqrs.Event{UserCreated{c.UserId, c.Name}}, nil
}

func (w *Winery) HandleAddVisit(c AddVisit) ([]cqrs.Event,error) {
	return []cqrs.Event{VisitAdded{c.UserId,c.time,c.WineryId,}}, nil
}




type CreateUser struct {
	UserId
	Name string
}
func (c CreateUser) TargetAggregateId() string {
	return string(c.UserId)
}

type AddVisit struct {
	UserId
	time string
	WineryId
}
func (c AddVisit) TargetAggregateId() string {
	return string(c.UserId)
}



type UserCreated struct {
	UserId
	Name string
}
func (c UserCreated) AggregateId() string {
	return string(c.UserId)
}
type VisitAdded struct {
	UserId
	time string
	WineryId
}
func (c VisitAdded) AggregateId() string {
	return string(c.UserId)
}
