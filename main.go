package main

import (
	"fmt"

	abs "github.com/heisenberg8055/go-patterns/structural/facade"
)

func main() {

	walletFacade := abs.NewWalletFacade("test", 123)
	err := walletFacade.AddMoneyToWallet("test", 123, 14)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = walletFacade.DeductMoneyFromWallet("test", 123, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
}
