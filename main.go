package main

import (
	"log"

	abs "github.com/heisenberg8055/go-patterns/behavioral/state"
)

func main() {
	vendingMachine := abs.NewVendingMachine(1, 10)

	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}

}
