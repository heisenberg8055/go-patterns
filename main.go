package main

import (
	"fmt"

	abs "github.com/heisenberg8055/go-patterns/structural/flyweight"
)

func main() {
	game := abs.NewGame()
	game.AddTerrorist(abs.TerroristDressType)
	game.AddTerrorist(abs.TerroristDressType)
	game.AddTerrorist(abs.TerroristDressType)
	game.AddTerrorist(abs.TerroristDressType)

	game.AddCounterTerrorist(abs.CounterTerroristDressType)
	game.AddCounterTerrorist(abs.CounterTerroristDressType)
	game.AddCounterTerrorist(abs.CounterTerroristDressType)
	game.AddCounterTerrorist(abs.CounterTerroristDressType)

	dressFactoryInstance := abs.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("Color: %s\tType: %s\n", dress.GetColor(), dressType)
	}
}
