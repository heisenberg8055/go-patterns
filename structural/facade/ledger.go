package facade

import "fmt"

type Ledger struct {
}

func (l *Ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Account ID: %v, Transaction Type: %v, Amount: %d", accountID, txnType, amount)
}
