// Package parseamount is exercise 3. Read CONCEPT.md first.
package parseamount

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseAmount converts a dollar string like "12.34" into cents (1234).
// It returns an error wrapping ErrEmpty or ErrInvalid on bad input.
func ParseAmount(s string) (int, error) {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return 0, ErrEmpty
	}

	invalid := fmt.Errorf("parsing %q: %w", s, ErrInvalid)

	dollarsPart, centsPart, hasDot := strings.Cut(trimmed, ".")

	dollars, ok := parseDigits(dollarsPart)
	if !ok {
		return 0, invalid
	}

	cents := 0
	if hasDot {
		if len(centsPart) < 1 || len(centsPart) > 2 {
			return 0, invalid
		}
		c, ok := parseDigits(centsPart)
		if !ok {
			return 0, invalid
		}
		if len(centsPart) == 1 {
			c *= 10 // "0.5" means 50 cents, not 5
		}
		cents = c
	}

	return dollars*100 + cents, nil
}

// parseDigits parses a non-empty string made only of ASCII digits.
// It rejects signs like "+" or "-", which strconv.Atoi would accept.
func parseDigits(s string) (int, bool) {
	if s == "" {
		return 0, false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return n, true
}
