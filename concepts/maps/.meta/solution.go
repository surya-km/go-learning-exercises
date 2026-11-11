// Package mapspractice practices the maps concept. Read CONCEPT.md first.
package mapspractice

import "slices"

// TotalByCategory returns total cents per category. Never returns nil.
func TotalByCategory(txns []Transaction) map[string]int {
	totals := map[string]int{}
	for _, t := range txns {
		totals[t.Category] += t.AmountCents
	}
	return totals
}

// OverBudget returns, sorted alphabetically, the categories whose spending is
// greater than their budget. Spending in a category with no budget entry counts
// as over budget. Never returns nil.
func OverBudget(spent, budget map[string]int) []string {
	over := []string{}
	for category, amount := range spent {
		limit, ok := budget[category]
		if (ok && amount > limit) || (!ok && amount > 0) {
			over = append(over, category)
		}
	}
	slices.Sort(over)
	return over
}
