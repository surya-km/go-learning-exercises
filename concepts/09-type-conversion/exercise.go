// Package typeconversion practices the type-conversion concept. Read CONCEPT.md first.
package typeconversion

// AverageCents returns totalCents / count as a float64, keeping the decimal part.
// It returns 0 if count <= 0.
func AverageCents(totalCents, count int) float64 {
	// TODO: implement
	return 0
}

// OrderSummary parses quantity and unitCents and returns a summary
// like "3 items, 1050 cents". It returns "" and an error for invalid numbers.
func OrderSummary(quantity, unitCents string) (string, error) {
	// TODO: implement
	return "", nil
}
