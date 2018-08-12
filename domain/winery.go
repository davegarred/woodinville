package domain

import "github.com/davegarred/cqrs"

type WineryId string

type Winery struct {
	wineryId WineryId
	Name string
	Address string
	City string
	Zip string
}

func (w *Winery) HandleCreateWinery(c CreateWinery) ([]cqrs.Event,error) {
	return []cqrs.Event{WineryCreated{c.WineryId}}, nil
}

func (w *Winery) HandleUpdateWineryAddress(c UpdateWineryAddress) ([]cqrs.Event,error) {
	event := WineryAddressUpdated{
		c.WineryId,
		c.Name,
		c.Address,
		c.City,
		c.Zip,
	}
	return []cqrs.Event{event}, nil
}



type CreateWinery struct {
	WineryId
}
func (c CreateWinery) TargetAggregateId() string {
	return string(c.WineryId)
}

type UpdateWineryAddress struct {
	WineryId
	Name string
	Address string
	City string
	Zip string
}
func (c UpdateWineryAddress) TargetAggregateId() string {
	return string(c.WineryId)
}


type WineryCreated struct {
	WineryId
}
func (c WineryCreated) AggregateId() string {
	return string(c.WineryId)
}

type WineryAddressUpdated struct {
	WineryId
	Name string
	Address string
	City string
	Zip string
}
func (c WineryAddressUpdated) AggregateId() string {
	return string(c.WineryId)
}
