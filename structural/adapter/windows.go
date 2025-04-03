package adapter

import "fmt"

type Windows struct {
}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("USB Connector is plugged into windows machine")
}
