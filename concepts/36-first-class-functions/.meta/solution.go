// Package firstclassfunctions practices the first-class functions concept. Read CONCEPT.md first.
package firstclassfunctions

// Filter returns the transactions for which keep returns true, in original order.
// It never returns nil and does not modify txns.
func Filter(txns []Transaction, keep func(Transaction) bool) []Transaction {
	result := []Transaction{}
	for _, t := range txns {
		if keep(t) {
			result = append(result, t)
		}
	}
	return result
}

// SpendingTracker returns a closure that subtracts each amount from a budget of
// limitCents. It returns the remaining budget and whether the amount was accepted.
// Amounts that would go over the budget are rejected and not subtracted.
func SpendingTracker(limitCents int) func(amountCents int) (remaining int, ok bool) {
	remaining := limitCents
	return func(amountCents int) (int, bool) {
		if amountCents > remaining {
			return remaining, false
		}
		remaining -= amountCents
		return remaining, true
	}
}
