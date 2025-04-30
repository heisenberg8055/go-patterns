package observer

import "fmt"

type Item struct {
	observerList map[Observer]bool
	Name         string
	InStock      bool
}

func NewItem(name string) *Item {
	return &Item{Name: name, observerList: make(map[Observer]bool)}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.Name)
	i.InStock = true
	i.notifyAll()
}

func (i *Item) Register(o Observer) {
	i.observerList[o] = true
}

func (i *Item) Deregister(o Observer) {
	delete(i.observerList, o)
}

func (i *Item) notifyAll() {
	for observer := range i.observerList {
		observer.update(i.Name)
	}
}
