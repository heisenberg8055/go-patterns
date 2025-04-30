package observer

import "fmt"

type Customer struct {
	Id string
}

func (c *Customer) getId() string {
	return c.Id
}

func (c *Customer) update(itemName string) {
	fmt.Printf("Sending emails to customer %s for item %s\n", c.Id, itemName)
}
