package builder

// Concrete Builder
type NormalBuilder struct {
	windowType string
	doorType   string
	floor      uint8
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) setWindowType() {
	n.windowType = "Wooden Window"
}

func (n *NormalBuilder) setDoorType() {
	n.doorType = "Wooden Door"
}

func (n *NormalBuilder) setNumFloor() {
	n.floor = 5
}

func (n *NormalBuilder) getHouse() House {
	return House{
		windowType: n.windowType,
		doorType:   n.doorType,
		floor:      n.floor,
	}
}
