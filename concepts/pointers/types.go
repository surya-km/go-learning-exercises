package pointers

import "errors"

// Account is a bank account. Money is stored in integer cents.
type Account struct {
	Owner        string
	BalanceCents int64
}

// Errors returned by Withdraw and Transfer.
var (
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrNilAccount        = errors.New("nil account")
)
