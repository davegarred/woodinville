package domain

type UserCreated struct {
	UserId
	Name string
}
func (c UserCreated) AggregateId() string {
	return string(c.UserId)
}

type UserIsSetAsAdmin struct {
	UserId
}
func (c UserIsSetAsAdmin) AggregateId() string {
	return string(c.UserId)
}

type UserIsSetAsNotAdmin struct {
	UserId
}
func (c UserIsSetAsNotAdmin) AggregateId() string {
	return string(c.UserId)
}

type VisitAdded struct {
	UserId
	Time string
	WineryId
}
func (c VisitAdded) AggregateId() string {
	return string(c.UserId)
}
