package main

import (
	"fmt"

	abs "github.com/heisenberg8055/go-patterns/behavioral/visitor"
)

func main() {
	square := &abs.Square{Side: 2}
	circle := &abs.Circle{Radius: 3}
	rectangle := &abs.Rectangle{L: 2, B: 3}

	areaCalculator := &abs.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &abs.MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)

}
