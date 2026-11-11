// Package typedefinitions practices the type definitions concept. Read CONCEPT.md first.
package typedefinitions

import "fmt"

// Format returns the amount as dollars, like "$12.34" or "-$12.50".
func (c Cents) Format() string {
	sign := ""
	if c < 0 {
		sign = "-"
		c = -c
	}
	return fmt.Sprintf("%s$%d.%02d", sign, c/100, c%100)
}

// Split divides the amount into n parts that add up to the original.
// Leftover cents go to the first parts. It returns nil if n <= 0.
func (c Cents) Split(n int) []Cents {
	if n <= 0 {
		return nil
	}
	base := c / Cents(n)
	leftover := int(c % Cents(n))
	parts := make([]Cents, n)
	for i := range parts {
		parts[i] = base
		if i < leftover {
			parts[i]++
		}
	}
	return parts
}

// TimesPerYear returns how many payments happen in a year (0 if unknown).
func (f Frequency) TimesPerYear() int {
	switch f {
	case Weekly:
		return 52
	case Monthly:
		return 12
	case Quarterly:
		return 4
	case Yearly:
		return 1
	default:
		return 0
	}
}

// YearlyCost returns the cost per year of paying amount at frequency f.
func YearlyCost(amount Cents, f Frequency) Cents {
	return amount * Cents(f.TimesPerYear())
}
