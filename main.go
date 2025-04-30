package main

import (
	abs "github.com/heisenberg8055/go-patterns/behavioral/observer"
)

func main() {
	shirtItem := abs.NewItem("Denim Shirt")

	obsFirst := &abs.Customer{Id: "test@gmail.com"}
	obsSecond := &abs.Customer{Id: "test1@gmail.com"}

	shirtItem.Register(obsFirst)
	shirtItem.Register(obsSecond)

	shirtItem.UpdateAvailability()
}
