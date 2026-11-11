// Package typedefinitions practices the type definitions concept. Read CONCEPT.md first.
package typedefinitions

// Format returns the amount as dollars, like "$12.34" or "-$12.50".
func (c Cents) Format() string {
	// TODO: implement
	return ""
}

// Split divides the amount into n parts that add up to the original.
// Leftover cents go to the first parts. It returns nil if n <= 0.
func (c Cents) Split(n int) []Cents {
	// TODO: implement
	return nil
}

// TimesPerYear returns how many payments happen in a year (0 if unknown).
func (f Frequency) TimesPerYear() int {
	// TODO: implement
	return 0
}

// YearlyCost returns the cost per year of paying amount at frequency f.
func YearlyCost(amount Cents, f Frequency) Cents {
	// TODO: implement
	return 0
}
