package storage

import "github.com/davegarred/woodinville/domain"

var locations map[domain.WineryId]*Location
var area []*Location
var users map[domain.UserId]*User

func init() {
	users = make(map[domain.UserId]*User)
	locations = make(map[domain.WineryId]*Location)
	area = make([]*Location, 0)

	addUser("MEL", "Melissa", false)
	addUser("JO", "Joanne", false)
	addUser("DAV", "Dave", true)
	addLocation(&Location{"DAR", 47.7322201,-122.14273, "Darby","14450 Redmond-Woodinville Rd NE","Woodinville", "98072"})
}

type Location struct {
	Id   domain.WineryId	`json:"id"`
	Lat float32 	`json:"lat"`
	Long float32	`json:"long"`
	Name string	`json:"name"`
	Address string `json:"address"`
	City string `json:"city"`
	Zip string `json:"zip"`
}

type User struct {
	Id domain.UserId `json:"id"`
	Name string `json:"name"`
	Admin bool `json:"admin"`
	Visits map[domain.WineryId][]string `json:"visits"`
}
func FindArea() []*Location {
	return area
}
func FindLocation(id domain.WineryId) *Location {
	return locations[id]
}

func FindUser(id domain.UserId) *User {
	return users[id]
}

func addUser(id domain.UserId, name string, admin bool) {
	users[id] = &User{id, name, admin,make(map[domain.WineryId][]string)}
}
func addLocation(l *Location) {
	locations[l.Id] = l
	area = append(area, l)
}