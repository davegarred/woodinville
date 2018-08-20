package domain

type WineryCreated struct {
	WineryId
	Name    string
}

func (c WineryCreated) AggregateId() string {
	return string(c.WineryId)
}

type WineryPositionUpdated struct {
	WineryId
	Lat float32
	Long float32
}
func (c WineryPositionUpdated) AggregateId() string {
	return string(c.WineryId)
}

type WineryAddressUpdated struct {
	WineryId
	Address string
	City    string
	Zip     string
}

func (c WineryAddressUpdated) AggregateId() string {
	return string(c.WineryId)
}
