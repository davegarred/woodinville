package storage

import (
	"github.com/davegarred/woodinville/domain"
	"github.com/davegarred/woodinville/query"
)

//var wineries []*WineryQuery
var users map[domain.UserId]*query.UserQuery
var locations map[domain.WineryId]*query.WineryQuery
var areaQuery *query.AreaQuery

var wineryIdentifierType IdentifierType
var wineryIdentifierFactory *IdentifierFactory

func init() {
	wineryIdentifierType = IdentifierType("WL")
	wineryIdentifierFactory = NewIdentifierFactory(wineryIdentifierType)

	users = make(map[domain.UserId]*query.UserQuery)
	locations = make(map[domain.WineryId]*query.WineryQuery)

	addUser("MEL", "Melissa", false)
	addUser("JO", "Joanne", false)
	addUser("DAV", "Dave", true)
	addLocation(&query.WineryQuery{"DAR", 47.7318,-122.14036, "Darby","14450 Redmond-Woodinville Rd NE","Woodinville", "98072", nil})
	areaQuery = &query.AreaQuery{"SEA", []domain.WineryId{"DAR"}, []query.RecommendedWinery{}}
}

func ResetTests() {
	locations = make(map[domain.WineryId]*query.WineryQuery)
	users = make(map[domain.UserId]*query.UserQuery)
	areaQuery.Wineries = []domain.WineryId{}
}

func NextWineryIdentifier() string {
	return wineryIdentifierFactory.Next()
}


func FindArea() *query.AreaQuery {
	return areaQuery
}
func UpdateArea(a *query.AreaQuery) {
	areaQuery = a
}

func FindWineries() []*query.WineryQuery {
	result := make([]*query.WineryQuery, len(locations))
	i := 0
	for _,l := range locations {
		result[i] = l
		i++
	}
	return result
}
func FindLocation(id domain.WineryId) *query.WineryQuery {
	return locations[id]
}
func UpdateLocation(id domain.WineryId, q *query.WineryQuery) {
	locations[id] = q
}

func FindUser(id domain.UserId) *query.UserQuery {
	return users[id]
}
func UpdateUser(id domain.UserId, q *query.UserQuery) {
	users[id] = q
}

func addUser(id domain.UserId, name string, admin bool) {
	users[id] = &query.UserQuery{id, name, admin,make(map[domain.WineryId][]string)}
}
func addLocation(l *query.WineryQuery) {
	locations[l.WineryId] = l
	//wineries = append(wineries, l)
}