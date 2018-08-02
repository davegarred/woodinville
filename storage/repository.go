package storage

var locations map[LocationId]*Location
var area []*Location
var users map[UserId]*User

func init() {
	users = make(map[UserId]*User)
	locations = make(map[LocationId]*Location)
	area = make([]*Location, 0)

	addUser("MEL", "Melissa")
	addUser("JO", "Joanne")
	addUser("DAV", "Dave")
	addLocation(&Location{"DAR", "Darby","14450 Redmond-Woodinville Rd NE","Woodinville", "98072"})
}

type LocationId string
type Location struct {
	Id   LocationId	`json:"id"`
	Name string	`json:"name"`
	Address string `json:"address"`
	City string `json:"city"`
	Zip string `json:"zip"`
}

type UserId string
type User struct {
	Id UserId `json:"id"`
	Name string `json:"name"`
	Visits map[LocationId][]string `json:"visits"`
}
func FindArea() []*Location {
	return area
}
func FindLocation(id LocationId) *Location {
	return locations[id]
}

func FindUser(id UserId) *User {
	return users[id]
}

func addUser(id UserId, name string) {
	users[id] = &User{id, name, make(map[LocationId][]string)}
}
func addLocation(l *Location) {
	locations[l.Id] = l
	area = append(area, l)
}