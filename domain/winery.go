package domain

import "github.com/davegarred/cqrs"

type WineryId string

type Winery struct {
	wineryId WineryId
	//Name string
	//Address string
	//City string
	//Zip string
}

func (w *Winery) HandleCreateWinery(c CreateWinery) ([]cqrs.Event, error) {
	return []cqrs.Event{WineryCreated{c.WineryId, c.Name}}, nil
}

func (w *Winery) HandleUpdateWineryPosition(c UpdateWineryPosition) ([]cqrs.Event, error) {
	return []cqrs.Event{WineryPositionUpdated{c.WineryId, c.Lat, c.Long}}, nil
}

func (w *Winery) HandleUpdateWineryAddress(c UpdateWineryAddress) ([]cqrs.Event, error) {
	event := WineryAddressUpdated{
		c.WineryId,
		c.Address,
		c.City,
		c.Zip,
	}
	return []cqrs.Event{event}, nil
}
