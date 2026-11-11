// Package booleans practices the booleans concept. Read CONCEPT.md first.
package booleans

// CanWithdraw reports whether amountCents can be withdrawn: the account must
// not be frozen, the amount must be positive and no more than balanceCents.
func CanWithdraw(balanceCents, amountCents int, frozen bool) bool {
	// TODO: implement
	return false
}

// NeedsReview reports whether a payment needs a manual fraud review: either it is
// 100000 cents or more, or it is international and going to a new payee.
func NeedsReview(amountCents int, international, newPayee bool) bool {
	// TODO: implement
	return false
}
