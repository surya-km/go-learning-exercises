package money

import "strconv"

// Cents is an amount of money in cents. It is exported (capital C), so other
// packages can use it as money.Cents.
type Cents int

// twoDigits formats n (0-99) with a leading zero, e.g. 5 -> "05".
// It is unexported (lowercase t): code in package money can call it, but other
// packages cannot.
func twoDigits(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}
