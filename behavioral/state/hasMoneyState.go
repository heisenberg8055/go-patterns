package state

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) requestItem() error {
	return fmt.Errorf("item Dispense in progress")
}

func (i *HasMoneyState) addItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *HasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item...")
	i.vendingMachine.itemCount -= 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.SetState(i.vendingMachine.hasItem)
	}
	return nil
}
