package factory

type Mp5 struct {
	Gun
}

func newMp5() IGun {
	return &Mp5{Gun: Gun{name: "MP5", power: 3}}
}
