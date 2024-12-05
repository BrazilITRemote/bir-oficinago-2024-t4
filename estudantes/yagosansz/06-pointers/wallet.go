package main

type Wallet struct {
	balance int
}

// IMPORTANT: Pointer dereferencing is to get the value stored in memory!
// We don't need to do `(*w).balance` (explicit dereference) because Go will
// automatically do that for us.
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Technically, this one would work with receiving a copy and not a pointer,
// but by convetion we should keep method receiver types the same for consistency!
func (w *Wallet) Balance() int {
	return w.balance
}
