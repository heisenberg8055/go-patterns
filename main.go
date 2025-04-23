package main

import (
	abs "github.com/heisenberg8055/go-patterns/behavioral/command"
)

func main() {
	tv := &abs.Tv{}

	onCommand := &abs.OnCommand{Device: tv}

	offCommand := &abs.OffCommand{Device: tv}

	onButton := &abs.Button{Command: onCommand}

	onButton.Press()

	offButton := &abs.Button{Command: offCommand}

	offButton.Press()
}
