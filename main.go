package main

import (
	abs "github.com/heisenberg8055/go-patterns/behavioral/strategy"
)

func main() {
	lfu := &abs.Lfu{}
	cache := abs.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")
	cache.Add("c", "3")

	lru := &abs.Lru{}

	cache.SetEvictionPolicy(lru)

	cache.Add("d", "4")

	fifo := &abs.Fifo{}
	cache.SetEvictionPolicy(fifo)

	cache.Add("e", "5")

}
