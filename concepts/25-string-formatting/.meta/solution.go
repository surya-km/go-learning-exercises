// Package stringformatting practices the string-formatting concept. Read CONCEPT.md first.
package stringformatting

import "fmt"

// StatementLine returns the description left-aligned in 20 characters, a space,
// and the amount like "$4.50" right-aligned in 10 characters.
func StatementLine(description string, cents int) string {
	sign := ""
	if cents < 0 {
		sign = "-"
		cents = -cents
	}
	amount := fmt.Sprintf("%s$%d.%02d", sign, cents/100, cents%100)
	return fmt.Sprintf("%-20s %10s", description, amount)
}

// SpendingShare returns a line like "Groceries: 25.0%" showing spentCents
// as a percentage of totalCents with one decimal place.
func SpendingShare(category string, spentCents, totalCents int) string {
	percent := 0.0
	if totalCents != 0 {
		percent = float64(spentCents) / float64(totalCents) * 100
	}
	return fmt.Sprintf("%s: %.1f%%", category, percent)
}
