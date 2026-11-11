// Package regularexpressions practices the regular-expressions concept. Read CONCEPT.md first.
package regularexpressions

import (
	"regexp"
	"strconv"
)

var (
	routingRe = regexp.MustCompile(`^\d{9}$`)
	amountRe  = regexp.MustCompile(`\$(\d+)(?:\.(\d{2}))?`)
)

// ValidRoutingNumber reports whether s is exactly 9 digits.
func ValidRoutingNumber(s string) bool {
	return routingRe.MatchString(s)
}

// ExtractAmounts finds every "$12" or "$12.34" style amount in text and returns
// each one in cents, in order.
func ExtractAmounts(text string) []int {
	var amounts []int
	for _, m := range amountRe.FindAllStringSubmatch(text, -1) {
		dollars, err := strconv.Atoi(m[1])
		if err != nil {
			continue // only possible for absurdly long digit runs
		}
		cents := 0
		if m[2] != "" {
			cents, _ = strconv.Atoi(m[2]) // always two digits, so this can't fail
		}
		amounts = append(amounts, dollars*100+cents)
	}
	return amounts
}
