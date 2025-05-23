package state

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) requestItem() error {
	return fmt.Errorf("item out of stock")
}

func (i *NoItemState) addItem(count int) error {
	i.vendingMachine.IncrementItemCount(count)
	i.vendingMachine.SetState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *NoItemState) dispenseItem() error {
	return fmt.Errorf("item out of stock")
}
