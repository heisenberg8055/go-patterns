package bridge

import "fmt"

type Epson struct {
}

func (e *Epson) PrintFile() {
	fmt.Println("Epson Working")
}
