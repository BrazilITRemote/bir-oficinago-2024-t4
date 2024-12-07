package main

import (
	"errors"
	"fmt"
)

// The var keyword allows us to define values global to the package!
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// IMPORTANT: Pointer dereferencing is to get the value stored in memory!
// We don't need to do `(*w).balance` (explicit dereference) because Go will
// automatically do that for us.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Technically, this one would work with receiving a copy and not a pointer,
// but by convetion we should keep method receiver types the same for consistency!
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
