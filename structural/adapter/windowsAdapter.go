package adapter

import "fmt"

type WindowsAdapter struct {
	WindowsMachine *Windows
}

func (wA *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts lightning signal to usb")
	wA.WindowsMachine.InsertIntoUSBPort()
}
