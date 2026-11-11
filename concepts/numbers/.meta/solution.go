// Package numbers practices the numbers concept. Read CONCEPT.md first.
package numbers

// SplitBill splits totalCents between people as evenly as possible.
// Earlier people pay the leftover cents. It returns nil if people <= 0.
func SplitBill(totalCents, people int) []int {
	if people <= 0 {
		return nil
	}

	base := totalCents / people
	leftover := totalCents % people

	shares := make([]int, people)
	for i := range shares {
		shares[i] = base
		if i < leftover {
			shares[i]++
		}
	}
	return shares
}

// MakeChange returns the fewest quarters, dimes, nickels and pennies
// that add up to cents. Negative amounts return all zeros.
func MakeChange(cents int) (quarters, dimes, nickels, pennies int) {
	if cents <= 0 {
		return 0, 0, 0, 0
	}

	quarters = cents / 25
	cents %= 25
	dimes = cents / 10
	cents %= 10
	nickels = cents / 5
	pennies = cents % 5
	return quarters, dimes, nickels, pennies
}
