package storage

import "github.com/davegarred/woodinville/domain"

var area []*WineryQuery
var users map[domain.UserId]*UserQuery
var locations map[domain.WineryId]*WineryQuery

func init() {
	users = make(map[domain.UserId]*UserQuery)
	locations = make(map[domain.WineryId]*WineryQuery)
	area = make([]*WineryQuery, 0)

	addUser("MEL", "Melissa", false)
	addUser("JO", "Joanne", false)
	addUser("DAV", "Dave", true)
	addLocation(&WineryQuery{"DAR", 47.7322201,-122.14273, "Darby","14450 Redmond-Woodinville Rd NE","Woodinville", "98072"})
}

//type Location struct {
//	Id   domain.WineryId	`json:"id"`
//	Lat float32 	`json:"lat"`
//	Long float32	`json:"long"`
//	Name string	`json:"name"`
//	Address string `json:"address"`
//	City string `json:"city"`
//	Zip string `json:"zip"`
//}

//type User struct {
//	Id domain.UserId `json:"id"`
//	Name string `json:"name"`
//	Admin bool `json:"admin"`
//	Visits map[domain.WineryId][]string `json:"visits"`
//}
func FindArea() []*WineryQuery {
	return area
}
func FindLocation(id domain.WineryId) *WineryQuery {
	return locations[id]
}

func FindUser(id domain.UserId) *UserQuery {
	return users[id]
}

func addUser(id domain.UserId, name string, admin bool) {
	users[id] = &UserQuery{id, name, admin,make(map[domain.WineryId][]string)}
}
func addLocation(l *WineryQuery) {
	locations[l.WineryId] = l
	area = append(area, l)
}