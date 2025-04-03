package main

import (
	abs "github.com/heisenberg8055/go-patterns/structural/bridge"
)

func main() {

	hpPrinter := &abs.Hp{}
	epsonPrinter := &abs.Epson{}

	mac := &abs.Mac{}
	mac.SetPrinter(hpPrinter)
	mac.Print()

	mac.SetPrinter(epsonPrinter)
	mac.Print()

	windows := &abs.Windows{}
	windows.SetPrinter(hpPrinter)
	windows.Print()

	windows.SetPrinter(epsonPrinter)
	windows.Print()

}
