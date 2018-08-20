package storage

var locations map[LocationId]*Location
var area []*Location
var users map[UserId]*User

func init() {
	users = make(map[UserId]*User)
	locations = make(map[LocationId]*Location)
	area = make([]*Location, 0)

	addUser("MEL", "Melissa", false)
	addUser("JO", "Joanne", false)
	addUser("DAV", "Dave", true)
	addLocation(&Location{"DAR", 47.7322201,-122.14273, "Darby","14450 Redmond-Woodinville Rd NE","Woodinville", "98072"})
}

type LocationId string
type Location struct {
	Id   LocationId	`json:"id"`
	Lat float32 	`json:"lat"`
	Long float32	`json:"long"`
	Name string	`json:"name"`
	Address string `json:"address"`
	City string `json:"city"`
	Zip string `json:"zip"`
}

type UserId string
type User struct {
	Id UserId `json:"id"`
	Name string `json:"name"`
	Admin bool `json:"admin"`
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

func addUser(id UserId, name string, admin bool) {
	users[id] = &User{id, name, admin,make(map[LocationId][]string)}
}
func addLocation(l *Location) {
	locations[l.Id] = l
	area = append(area, l)
}