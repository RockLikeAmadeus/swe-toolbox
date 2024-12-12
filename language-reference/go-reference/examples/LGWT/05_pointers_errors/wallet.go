package wallet

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

// New type based on existing type
type Bitcoin int

// Can still define methods
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// The receiver is a pointer to a Wallet, so that we can modify the instance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Using a pointer receiver here isn't necessary, but should keep our receiver types consistent
func (w *Wallet) Balance() Bitcoin {
	// `return (*w).balance` is also valid, but struct pointers are automatically de-referenced
	return w.balance
}

// Errors are values
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
