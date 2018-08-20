package domain

type CreateUser struct {
	UserId
	Name string
}

func (c CreateUser) TargetAggregateId() string {
	return string(c.UserId)
}
type SetUserAdmin struct {
	UserId
	IsAdmin bool
}
func (c SetUserAdmin) TargetAggregateId() string {
	return string(c.UserId)
}

type AddVisit struct {
	UserId
	Time string
	WineryId
}
func (c AddVisit) TargetAggregateId() string {
	return string(c.UserId)
}
