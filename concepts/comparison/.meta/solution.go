// Package comparison practices the comparison concept. Read CONCEPT.md first.
package comparison

// CompareAmounts returns -1 if aCents < bCents, 0 if they are equal,
// and 1 if aCents > bCents.
func CompareAmounts(aCents, bCents int) int {
	if aCents < bCents {
		return -1
	}
	if aCents > bCents {
		return 1
	}
	return 0
}

// EarlierDate returns whichever of two "YYYY-MM-DD" dates comes first.
// If they are the same, it returns a.
func EarlierDate(a, b string) string {
	if b < a {
		return b
	}
	return a
}
