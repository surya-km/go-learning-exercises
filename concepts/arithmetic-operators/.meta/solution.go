// Package arithmeticoperators practices the arithmetic operators concept. Read CONCEPT.md first.
package arithmeticoperators

// TipCents returns percent% of billCents, rounded to the nearest cent
// (exactly half a cent rounds up).
func TipCents(billCents, percent int) int {
	return (billCents*percent + 50) / 100
}

// RoundUp rounds cents up to the next whole dollar. It returns the rounded
// amount and the spare change (the difference) that would be saved.
func RoundUp(cents int) (roundedCents, spareChangeCents int) {
	remainder := cents % 100
	if remainder == 0 {
		return cents, 0
	}
	spare := 100 - remainder
	return cents + spare, spare
}
