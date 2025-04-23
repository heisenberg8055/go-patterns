package mediator

import "fmt"

type FreightTrain struct {
	Mediator Mediator
}

func (g *FreightTrain) Arrive() {
	if !g.Mediator.canArrive(g) {
		fmt.Println("FreightTrain: Arrival Blocked!")
		return
	}
	fmt.Println("FreightsTrain: Arrived!")
}

func (g *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Departed!")
	g.Mediator.notifyAboutDeparture()
}

func (g *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Freight Train Arrival Permitted!")
	g.Arrive()
}
