// Package slicesmaps is exercise 2. Read CONCEPT.md first.
package slicesmaps

// Sum returns the total of all amounts.
func Sum(amounts []int) int {
	total := 0
	for _, amt := range amounts {
		total += amt
	}
	return total
}

// Above returns the amounts strictly greater than min, in original order.
func Above(amounts []int, min int) []int {
	result := []int{}
	for _, amt := range amounts {
		if amt > min {
			result = append(result, amt)
		}
	}
	return result
}

// TotalByCategory returns total cents per category. Never returns nil.
func TotalByCategory(txns []Transaction) map[string]int {
	totals := map[string]int{}
	for _, t := range txns {
		totals[t.Category] += t.AmountCents
	}
	return totals
}
