// Package slicespractice practices the slices concept. Read CONCEPT.md first.
package slicespractice

// Sum returns the total of all amounts.
func Sum(amounts []int) int {
	total := 0
	for _, amt := range amounts {
		total += amt
	}
	return total
}

// Above returns a new slice of the amounts strictly greater than min, in
// original order. It never returns nil and does not modify amounts.
func Above(amounts []int, min int) []int {
	result := []int{}
	for _, amt := range amounts {
		if amt > min {
			result = append(result, amt)
		}
	}
	return result
}

// Pages splits items into pages of pageSize items; the last page may be shorter.
// Each page is an independent copy. A pageSize of 0 or less, or no items,
// returns an empty, non-nil slice.
func Pages(items []string, pageSize int) [][]string {
	pages := [][]string{}
	if pageSize <= 0 {
		return pages
	}
	for start := 0; start < len(items); start += pageSize {
		end := min(start+pageSize, len(items))
		page := make([]string, end-start)
		copy(page, items[start:end])
		pages = append(pages, page)
	}
	return pages
}
