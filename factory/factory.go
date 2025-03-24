package factory

import "fmt"

func Factory() {
	ak47, _ := getGun("ak47")
	mp5, _ := getGun("mp5")

	printGun(ak47)
	printGun(mp5)
}

func printGun(g IGun) {
	fmt.Printf("Gun:\tName:%s\tPower:%d\n", g.getName(), g.getPower())
}
