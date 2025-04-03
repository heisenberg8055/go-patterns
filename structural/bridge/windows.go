package bridge

import "fmt"

type Windows struct {
	P Printer
}

func (m *Windows) Print() {
	fmt.Println("Print request for mac")
	m.P.PrintFile()
}

func (m *Windows) SetPrinter(p Printer) {
	m.P = p
}
