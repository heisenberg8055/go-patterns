package factory

type IGun interface {
	setName(name string)
	setPower(power uint8)
	getName() string
	getPower() uint8
}
