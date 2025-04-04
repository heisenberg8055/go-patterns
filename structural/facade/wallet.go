package facade

import "fmt"

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("wallet balance credited")
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance-amount < 0 {
		return fmt.Errorf("insufficient Balance")
	}
	w.balance -= amount
	fmt.Println("wallet balance debited")
	return nil
}
