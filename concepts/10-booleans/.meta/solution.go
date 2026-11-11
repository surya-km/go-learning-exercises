// Package booleans practices the booleans concept. Read CONCEPT.md first.
package booleans

// CanWithdraw reports whether amountCents can be withdrawn: the account must
// not be frozen, the amount must be positive and no more than balanceCents.
func CanWithdraw(balanceCents, amountCents int, frozen bool) bool {
	return !frozen && amountCents > 0 && amountCents <= balanceCents
}

// NeedsReview reports whether a payment needs a manual fraud review: either it is
// 100000 cents or more, or it is international and going to a new payee.
func NeedsReview(amountCents int, international, newPayee bool) bool {
	return amountCents >= 100000 || (international && newPayee)
}
