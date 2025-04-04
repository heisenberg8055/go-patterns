package facade

import "fmt"

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Creating new account")
	walletFacade := &WalletFacade{account: newAccount(accountID), securityCode: newSecurityCode(code), wallet: newWallet(), notification: &Notification{}, ledger: &Ledger{}}
	fmt.Println("Account Created")
	return walletFacade
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Started adding money to account")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return nil
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Started deducting money to account")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return nil
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletdebitNotification()
	w.ledger.makeEntry(accountID, "debit", amount)
	return nil
}
