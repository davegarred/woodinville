package storage

import (
	"github.com/davegarred/woodinville/domain"
	"github.com/davegarred/cqrs"
)

var wineries []*WineryQuery
var users map[domain.UserId]*UserQuery
var locations map[domain.WineryId]*WineryQuery
var areaQuery *AreaQuery
var eventStore cqrs.EventStore
var commandGateway *cqrs.CommandGateway

var wineryIdentifierType IdentifierType
var wineryIdentifierFactory *IdentifierFactory

func init() {
	wineryIdentifierType = IdentifierType("WL")
	wineryIdentifierFactory = NewIdentifierFactory(wineryIdentifierType)

	eventStore = cqrs.NewMemEventStore()
	commandGateway = cqrs.NewCommandGateway(eventStore)
	commandGateway.RegisterAggregate(&domain.User{})
	commandGateway.RegisterAggregate(&domain.Winery{})
	commandGateway.RegisterAggregate(&domain.Area{})
	commandGateway.RegisterQueryEventHandlers(&UserQueryEventListener{})
	commandGateway.RegisterQueryEventHandlers(&WineryQueryEventListener{})
	commandGateway.RegisterQueryEventHandlers(&AreaQueryEventListener{})

	users = make(map[domain.UserId]*UserQuery)
	locations = make(map[domain.WineryId]*WineryQuery)

	addUser("MEL", "Melissa", false)
	addUser("JO", "Joanne", false)
	addUser("DAV", "Dave", true)
	addLocation(&WineryQuery{"DAR", 47.7318,-122.14036, "Darby","14450 Redmond-Woodinville Rd NE","Woodinville", "98072", nil})
	areaQuery = &AreaQuery{"SEA", []domain.WineryId{"DAR"}, []RecommendedWinery{}}
}

func NextWineryIdentifier() string {
	return wineryIdentifierFactory.Next()
}
func Dispatch(command cqrs.Command) {
	commandGateway.Dispatch(command)
}

func FindArea() *AreaQuery {
	return areaQuery
}
func UpdateArea(a *AreaQuery) {
	areaQuery = a
}

func FindWineries() []*WineryQuery {
	result := make([]*WineryQuery, len(locations))
	i := 0
	for _,l := range locations {
		result[i] = l
		i++
	}
	return result
}
func FindLocation(id domain.WineryId) *WineryQuery {
	return locations[id]
}
func UpdateLocation(id domain.WineryId, q *WineryQuery) {
	locations[id] = q
}

func FindUser(id domain.UserId) *UserQuery {
	return users[id]
}
func UpdateUser(id domain.UserId, q *UserQuery) {
	users[id] = q
}

func addUser(id domain.UserId, name string, admin bool) {
	users[id] = &UserQuery{id, name, admin,make(map[domain.WineryId][]string)}
}
func addLocation(l *WineryQuery) {
	locations[l.WineryId] = l
	wineries = append(wineries, l)
}