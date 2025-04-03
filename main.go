package main

import (
	abs "github.com/heisenberg8055/go-patterns/structural/adapter"
)

func main() {
	client := &abs.Client{}
	mac := &abs.Mac{}
	client.InsertLightningConnectorIntoPC(mac)

	windowsMachine := &abs.Windows{}
	windowsAdapter := &abs.WindowsAdapter{WindowsMachine: windowsMachine}
	windowsAdapter.InsertIntoLightningPort()

}
