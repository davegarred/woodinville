package domain

type AreaWineryRecommended struct {
	AreaId
	Name    string
	Address string
	City    string
	Zip     string
}

func (c AreaWineryRecommended) AggregateId() string {
	return string(c.AreaId)
}
