package main

import (
	"fmt"

	abs "github.com/heisenberg8055/go-patterns/creational/singleton"
)

func main() {
	for i := 0; i < 30; i++ {
		go abs.GetInstance()
	}
	fmt.Scanln()
}
