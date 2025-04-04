package facade

import "fmt"

type Notification struct {
}

func (n *Notification) sendWalletCreditNotification() {
	fmt.Println("sending wallet credit notification")
}

func (n *Notification) sendWalletdebitNotification() {
	fmt.Println("sending wallet debit notification")
}
