package state

import "fmt"

type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State

	currentState State

	itemCount int
	itemPrice int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{itemCount: itemCount, itemPrice: itemPrice}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}

	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}

	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}

	noItemState := &NoItemState{
		vendingMachine: v,
	}

	v.SetState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.requestItem()

}

func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *VendingMachine) IncrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}
