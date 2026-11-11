// Package pointers practices the pointers concept. Read CONCEPT.md first.
package pointers

// Deposit adds cents to the balance. Amounts of zero or less are ignored.
func (a *Account) Deposit(cents int64) {
	if cents <= 0 {
		return
	}
	a.BalanceCents += cents
}

// Withdraw takes cents out of the balance. It returns ErrInsufficientFunds
// (leaving the balance unchanged) if there is not enough money.
// Amounts of zero or less are ignored.
func (a *Account) Withdraw(cents int64) error {
	if cents <= 0 {
		return nil
	}
	if cents > a.BalanceCents {
		return ErrInsufficientFunds
	}
	a.BalanceCents -= cents
	return nil
}

// Transfer moves cents from one account to another. It returns ErrNilAccount
// if either account is nil and ErrInsufficientFunds if from cannot cover it.
// On error, neither account changes.
func Transfer(from, to *Account, cents int64) error {
	if from == nil || to == nil {
		return ErrNilAccount
	}
	if err := from.Withdraw(cents); err != nil {
		return err
	}
	to.Deposit(cents)
	return nil
}
