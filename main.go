package main

import (
	"fmt"

	abs "github.com/heisenberg8055/go-patterns/structural/decorator"
)

func main() {

	normalPizza := &abs.VeggieMania{}

	pizzaCheese := &abs.CheeseTopping{Pizza: normalPizza}

	pizzaTomato := &abs.TomatoToppings{Pizza: normalPizza}

	pizzaBoth := &abs.TomatoToppings{Pizza: pizzaCheese}

	fmt.Printf("Normal: %d, Cheese: %d, Tomato: %d, Both: %d", normalPizza.GetPrice(), pizzaCheese.GetPrice(), pizzaTomato.GetPrice(), pizzaBoth.GetPrice())
}
