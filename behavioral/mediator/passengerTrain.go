package mediator

import "fmt"

type PassengerTrain struct {
	Mediator Mediator
}

func (g *PassengerTrain) Arrive() {
	if !g.Mediator.canArrive(g) {
		fmt.Println("PassengerTrain: Arrival Blocked!")
		return
	}
	fmt.Println("PassengerTrain: Arrived!")
}

func (g *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Departed!")
	g.Mediator.notifyAboutDeparture()
}

func (g *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Passenger Train Arrival Permitted!")
	g.Arrive()
}
