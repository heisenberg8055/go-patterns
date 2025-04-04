package flyweight

type Player struct {
	dress      Dress
	playerType string
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := GetDressFactorySingleInstance().getDressByType(dressType)
	return &Player{dress: dress, playerType: playerType}
}
