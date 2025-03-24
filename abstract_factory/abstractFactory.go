package abstract_factory

import "fmt"

func Abstract_factory() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoe(nikeShoe)
	printShoe(adidasShoe)
	printShirt(nikeShirt)
	printShirt(adidasShirt)
}

func printShoe(s IShoe) {
	fmt.Printf("Shoe:\tLogo: %s\tSize: %d\n", s.getLogo(), s.getSize())
}

func printShirt(s IShirt) {
	fmt.Printf("Shirt:\tLogo: %s\tSize: %d\n", s.getLogo(), s.getSize())
}
