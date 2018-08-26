package domain

type RecommendWinery struct {
	AreaId
	Name    string
	Address string
	City    string
	Zip     string
}

func (c RecommendWinery) TargetAggregateId() string {
	return string(c.AreaId)
}
