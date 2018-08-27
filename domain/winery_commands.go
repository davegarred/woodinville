package domain

type CreateWinery struct {
	WineryId
	Name    string
}

func (c CreateWinery) TargetAggregateId() string {
	return string(c.WineryId)
}

type UpdateWineryPosition struct {
	WineryId
	Lat float32
	Long float32
}

func (c UpdateWineryPosition) TargetAggregateId() string {
	return string(c.WineryId)
}

type UpdateWineryAddress struct {
	WineryId
	Address string
	City    string
	Zip     string
}

func (c UpdateWineryAddress) TargetAggregateId() string {
	return string(c.WineryId)
}

type AddWineryWithAddress struct {
	WineryId
	Name    string
	Address string
	City    string
	Zip     string
}

func (c AddWineryWithAddress) TargetAggregateId() string {
	return string(c.WineryId)
}
