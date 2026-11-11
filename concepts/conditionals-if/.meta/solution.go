// Package conditionalsif practices the conditionals-if concept. Read CONCEPT.md first.
package conditionalsif

// OverdraftFee returns the fee charged for a balance: nothing when the balance is
// zero or above or overdrawn by 500 cents or less, otherwise 2500 cents.
func OverdraftFee(balanceCents int) int {
	if balanceCents >= -500 {
		return 0
	}
	return 2500
}

// BudgetStatus returns "no budget", "over", "warning" or "ok" depending on
// how spentCents compares with budgetCents.
func BudgetStatus(spentCents, budgetCents int) string {
	if budgetCents <= 0 {
		return "no budget"
	}
	if spentCents > budgetCents {
		return "over"
	}
	if spentCents*100 >= budgetCents*80 {
		return "warning"
	}
	return "ok"
}
