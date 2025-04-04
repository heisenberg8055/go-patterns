package decorator

type VeggieMania struct {
}

func (p *VeggieMania) GetPrice() uint {
	return 15
}
