package strategy

import "fmt"

type Lfu struct {
}

func (l *Lfu) Evict(c *Cache) {
	fmt.Println("LFU")
}
