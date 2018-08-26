package domain

import "github.com/davegarred/cqrs"

type AreaId string

type Area struct {
	areaId AreaId
}

func (w *Area) HandleRecommendWinery(c RecommendWinery) ([]cqrs.Event, error) {
	return []cqrs.Event{AreaWineryRecommended{c.AreaId, c.Name, c.Address, c.City, c.Zip}}, nil
}
