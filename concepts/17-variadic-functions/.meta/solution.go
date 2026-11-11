// Package variadicfunctions practices the variadic functions concept. Read CONCEPT.md first.
package variadicfunctions

// MaxAmount returns the largest of amounts and true.
// With no amounts it returns 0 and false.
func MaxAmount(amounts ...int) (int, bool) {
	if len(amounts) == 0 {
		return 0, false
	}
	largest := amounts[0]
	for _, a := range amounts[1:] {
		if a > largest {
			largest = a
		}
	}
	return largest, true
}

// AddToBill returns billCents plus the sum of every item price.
func AddToBill(billCents int, items ...int) int {
	for _, item := range items {
		billCents += item
	}
	return billCents
}
