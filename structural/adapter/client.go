package adapter

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoPC(comp Computer) {
	fmt.Println("client inserts lightning connector into pc")
	comp.InsertIntoLightningPort()
}
