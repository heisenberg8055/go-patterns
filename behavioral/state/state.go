package state

type State interface {
	addItem(count int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}
