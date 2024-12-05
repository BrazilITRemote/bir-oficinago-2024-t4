package main

import "fmt"

type Wallet struct {
	balance int
}

// w Wallet is a copy of the original receiver!
// In our case the "original" receiver is coming from tests.
func (w Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
