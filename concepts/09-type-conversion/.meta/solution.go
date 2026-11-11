// Package typeconversion practices the type-conversion concept. Read CONCEPT.md first.
package typeconversion

import "strconv"

// AverageCents returns totalCents / count as a float64, keeping the decimal part.
// It returns 0 if count <= 0.
func AverageCents(totalCents, count int) float64 {
	if count <= 0 {
		return 0
	}
	return float64(totalCents) / float64(count)
}

// OrderSummary parses quantity and unitCents and returns a summary
// like "3 items, 1050 cents". It returns "" and an error for invalid numbers.
func OrderSummary(quantity, unitCents string) (string, error) {
	qty, err := strconv.Atoi(quantity)
	if err != nil {
		return "", err
	}
	unit, err := strconv.Atoi(unitCents)
	if err != nil {
		return "", err
	}

	noun := "items"
	if qty == 1 {
		noun = "item"
	}
	return strconv.Itoa(qty) + " " + noun + ", " + strconv.Itoa(qty*unit) + " cents", nil
}
