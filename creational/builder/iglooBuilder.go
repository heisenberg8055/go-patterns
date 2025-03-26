package builder

// Concete Builder

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      uint8
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (i *IglooBuilder) setWindowType() {
	i.windowType = "Snow Window"
}

func (i *IglooBuilder) setDoorType() {
	i.doorType = "Snow Door"
}

func (i *IglooBuilder) setNumFloor() {
	i.floor = 1
}

func (i *IglooBuilder) getHouse() House {
	return House{
		windowType: i.windowType,
		doorType:   i.doorType,
		floor:      i.floor,
	}
}
