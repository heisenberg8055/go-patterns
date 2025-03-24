package factory

type Gun struct {
	name  string
	power uint8
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power uint8) {
	g.power = power
}

func (g *Gun) getPower() uint8 {
	return g.power
}
