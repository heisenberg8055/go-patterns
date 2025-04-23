package command

import "fmt"

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning TV on!")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning TV off!")
}
