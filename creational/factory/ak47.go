package factory

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{Gun: Gun{name: "AK47", power: 4}}
}
