package flyweight

type Game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func NewGame() *Game {
	return &Game{terrorists: make([]*Player, 1), counterTerrorists: make([]*Player, 1)}
}

func (g *Game) AddTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	g.terrorists = append(g.terrorists, player)
}

func (g *Game) AddCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	g.counterTerrorists = append(g.counterTerrorists, player)
}
