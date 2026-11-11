// Package variables practices the variables concept. Read CONCEPT.md first.
package variables

// MonthEnd returns the balance at the end of the month and the total spent
// on rent and groceries.
func MonthEnd(startCents, incomeCents, rentCents, groceriesCents int) (endCents, spentCents int) {
	spentCents = rentCents + groceriesCents
	endCents = startCents + incomeCents - spentCents
	return endCents, spentCents
}

// Transfer moves up to amountCents from one account to another, never more
// than fromCents holds, and returns both new balances.
func Transfer(fromCents, toCents, amountCents int) (newFrom, newTo int) {
	moved := amountCents
	if moved > fromCents {
		moved = fromCents
	}
	if moved < 0 {
		moved = 0
	}

	newFrom = fromCents - moved
	newTo = toCents + moved
	return newFrom, newTo
}
