package strategy

import "fmt"

type Fifo struct {
}

func (f *Fifo) Evict(c *Cache) {
	fmt.Println("FIFO")
}
