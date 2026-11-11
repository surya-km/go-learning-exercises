// Package multiplereturnvalues practices the multiple return values concept. Read CONCEPT.md first.
package multiplereturnvalues

// SplitBill splits totalCents evenly between people.
// It returns the amount each person pays and the leftover cents.
// If people <= 0, it returns 0 and totalCents.
func SplitBill(totalCents, people int) (each, leftover int) {
	// TODO: implement
	return 0, 0
}

// CoinChange returns the fewest quarters, dimes, nickels and pennies that add up
// to cents. It returns all zeros when cents <= 0.
func CoinChange(cents int) (quarters, dimes, nickels, pennies int) {
	// TODO: implement
	return 0, 0, 0, 0
}
