package flyweight

import "fmt"

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		DressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	DressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.DressMap[dressType] != nil {
		return d.DressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.DressMap[dressType] = newTerroristDress()
	}
	if dressType == CounterTerroristDressType {
		d.DressMap[dressType] = newCounterTerroristDress()
	}
	return nil, fmt.Errorf("wrong Dress type passed")
}

func GetDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}
