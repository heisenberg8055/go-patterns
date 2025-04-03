package bridge

import "fmt"

type Hp struct {
}

func (e *Hp) PrintFile() {
	fmt.Println("Hp Working")
}
