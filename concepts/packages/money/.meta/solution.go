// Package money practices the packages concept (the package being imported).
// Read ../CONCEPT.md first.
package money

import "strconv"

// Format returns c as a dollar string like "$12.34" or "-$0.05".
func Format(c Cents) string {
	sign := ""
	if c < 0 {
		sign = "-"
		c = -c
	}
	dollars, cents := int(c)/100, int(c)%100
	return sign + "$" + strconv.Itoa(dollars) + "." + twoDigits(cents)
}
