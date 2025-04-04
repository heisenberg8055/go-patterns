package facade

import "fmt"

type Account struct {
	id string
}

func newAccount(accountId string) *Account {
	return &Account{id: accountId}
}

func (a *Account) checkAccount(accountId string) error {
	if a.id != accountId {
		return fmt.Errorf("wrong account id")
	}
	fmt.Printf("Account Verified: %v", a.id)
	return nil
}
