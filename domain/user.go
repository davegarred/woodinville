package domain

import "github.com/davegarred/cqrs"

type UserId string


type User struct {
	userId UserId
	//name   string
	//visits []WineryVisit
}

type WineryVisit struct {
	//WineryId
	//Time string
}

func (w *Winery) HandleCreateUser(c CreateUser) ([]cqrs.Event, error) {
	return []cqrs.Event{UserCreated{c.UserId, c.Name}}, nil
}

func (w *Winery) HandleSetUserAdmin(c SetUserAdmin) ([]cqrs.Event, error) {
	if c.IsAdmin {
		return []cqrs.Event{UserIsSetAsAdmin{c.UserId}}, nil
	}
	return []cqrs.Event{UserIsSetAsNotAdmin{c.UserId}}, nil
}

func (w *Winery) HandleAddVisit(c AddVisit) ([]cqrs.Event, error) {
	return []cqrs.Event{VisitAdded{c.UserId, c.Time, c.WineryId,}}, nil
}

