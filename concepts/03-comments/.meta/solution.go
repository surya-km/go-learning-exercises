// Package comments practices the comments concept. Read CONCEPT.md first.
package comments

import "fmt"

// CurrencySymbol is placed in front of every formatted amount.
const CurrencySymbol = "$"

// Transaction is one payment in or out of an account.
// AmountCents is negative for money going out.
type Transaction struct {
	Description string
	AmountCents int
}

// FormatAmount formats cents as a dollar string such as "$12.34" or "-$3.50".
func FormatAmount(cents int) string {
	sign := ""
	if cents < 0 {
		sign = "-"
		cents = -cents
	}
	return fmt.Sprintf("%s%s%d.%02d", sign, CurrencySymbol, cents/100, cents%100)
}

// Describe returns a one-line summary of t, such as "Coffee: $3.50".
func Describe(t Transaction) string {
	return t.Description + ": " + FormatAmount(t.AmountCents)
}
