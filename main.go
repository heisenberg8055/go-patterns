package main

import (
	"fmt"

	abs "github.com/heisenberg8055/go-patterns/behavioral/memento"
)

func main() {
	careTaker := abs.Caretaker{MementoArray: make([]*abs.Memento, 0)}

	originator := abs.Originator{State: "A"}
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	careTaker.AddMemento(originator.CreateMemento())

	originator.SetState("B")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	careTaker.AddMemento(originator.CreateMemento())

	originator.SetState("C")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	careTaker.AddMemento(originator.CreateMemento())

	originator.RestoreMemento(careTaker.GetMemento(1))
	fmt.Printf("Originator Current State: %s\n", originator.GetState())

	originator.RestoreMemento(careTaker.GetMemento(0))
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
}
