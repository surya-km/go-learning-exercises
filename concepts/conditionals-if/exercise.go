// Package conditionalsif practices the conditionals-if concept. Read CONCEPT.md first.
package conditionalsif

// OverdraftFee returns the fee charged for a balance: nothing when the balance is
// zero or above or overdrawn by 500 cents or less, otherwise 2500 cents.
func OverdraftFee(balanceCents int) int {
	// TODO: implement
	return 0
}

// BudgetStatus returns "no budget", "over", "warning" or "ok" depending on
// how spentCents compares with budgetCents.
func BudgetStatus(spentCents, budgetCents int) string {
	// TODO: implement
	return ""
}
