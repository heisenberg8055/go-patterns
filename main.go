package main

import (
	abs "github.com/heisenberg8055/go-patterns/behavioral/mediator"
)

func main() {
	stationMaster := abs.NewStationMaster()

	passengerTrain := &abs.PassengerTrain{Mediator: stationMaster}

	FreightTrain := &abs.FreightTrain{Mediator: stationMaster}

	passengerTrain.Arrive()
	FreightTrain.Arrive()
	passengerTrain.Depart()
}
