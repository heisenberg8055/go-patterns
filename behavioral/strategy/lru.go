package strategy

import "fmt"

type Lru struct {
}

func (l *Lru) Evict(c *Cache) {
	fmt.Println("LRU")
}
