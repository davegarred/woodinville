package storage

import (
	"github.com/davegarred/woodinville/domain"
	"github.com/davegarred/cqrs"
)

var area []*WineryQuery
var users map[domain.UserId]*UserQuery
var locations map[domain.WineryId]*WineryQuery
var eventStore cqrs.EventStore
var commandGateway *cqrs.CommandGateway

func init() {
	eventStore = cqrs.NewMemEventStore()
	commandGateway = cqrs.NewCommandGateway(eventStore)
	commandGateway.RegisterAggregate(&domain.User{})
	commandGateway.RegisterAggregate(&domain.Winery{})
	commandGateway.RegisterQueryEventHandlers(&UserWineryQueryEventListener{})

	users = make(map[domain.UserId]*UserQuery)
	locations = make(map[domain.WineryId]*WineryQuery)
	area = make([]*WineryQuery, 0)

	addUser("MEL", "Melissa", false)
	addUser("JO", "Joanne", false)
	addUser("DAV", "Dave", true)
	addLocation(&WineryQuery{"DAR", 47.7322201,-122.14273, "Darby","14450 Redmond-Woodinville Rd NE","Woodinville", "98072"})
}

func Dispatch(command cqrs.Command) {
	commandGateway.Dispatch(command)
}
func FindArea() []*WineryQuery {
	return area
}
func FindLocation(id domain.WineryId) *WineryQuery {
	return locations[id]
}

func FindUser(id domain.UserId) *UserQuery {
	return users[id]
}
func UpdateLocation(id domain.WineryId, q *WineryQuery) {
	locations[id] = q
}

func UpdateUser(id domain.UserId, q *UserQuery) {
	users[id] = q
}

func addUser(id domain.UserId, name string, admin bool) {
	users[id] = &UserQuery{id, name, admin,make(map[domain.WineryId][]string)}
}
func addLocation(l *WineryQuery) {
	locations[l.WineryId] = l
	area = append(area, l)
}