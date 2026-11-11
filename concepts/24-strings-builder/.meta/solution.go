// Package stringsbuilder practices the strings-builder concept. Read CONCEPT.md first.
package stringsbuilder

import (
	"strconv"
	"strings"
)

// Receipt returns one "Name: cents" line per item followed by a "TOTAL: sum" line.
// Every line ends with "\n".
func Receipt(items []LineItem) string {
	var sb strings.Builder
	total := 0
	for _, item := range items {
		sb.WriteString(item.Name)
		sb.WriteString(": ")
		sb.WriteString(strconv.Itoa(item.Cents))
		sb.WriteByte('\n')
		total += item.Cents
	}
	sb.WriteString("TOTAL: ")
	sb.WriteString(strconv.Itoa(total))
	sb.WriteByte('\n')
	return sb.String()
}

// GroupDigits formats cents as dollars with comma separators, like "1,234.50".
func GroupDigits(cents int) string {
	var sb strings.Builder
	if cents < 0 {
		sb.WriteByte('-')
		cents = -cents
	}

	digits := strconv.Itoa(cents / 100)
	for i := 0; i < len(digits); i++ {
		// Add a comma when the number of digits left is a multiple of 3.
		if i > 0 && (len(digits)-i)%3 == 0 {
			sb.WriteByte(',')
		}
		sb.WriteByte(digits[i])
	}

	rem := cents % 100
	sb.WriteByte('.')
	if rem < 10 {
		sb.WriteByte('0')
	}
	sb.WriteString(strconv.Itoa(rem))
	return sb.String()
}
