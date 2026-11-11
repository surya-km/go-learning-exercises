// Package floatingpointnumbers practices the floating-point-numbers concept. Read CONCEPT.md first.
package floatingpointnumbers

import "math"

// ApplyRate returns ratePercent percent of cents, rounded to the nearest cent.
// For example, ApplyRate(1999, 8.25) is 165.
func ApplyRate(cents int, ratePercent float64) int {
	exact := float64(cents) * ratePercent / 100
	return int(math.Round(exact))
}

// AlmostEqual reports whether a and b differ by at most tolerance.
func AlmostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
