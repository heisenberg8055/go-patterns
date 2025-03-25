package builder

import "fmt"

// Client Code
func Builder() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()
	printHouse(normalHouse)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()
	printHouse(iglooHouse)
}

func printHouse(house House) {
	fmt.Printf("%s\t%s\t%d\n", house.doorType, house.windowType, house.floor)
}
