// Package methods practices the methods concept. Read CONCEPT.md first.
package methods

import "fmt"

// IsDueBy reports whether the bill is due on or before day.
func (b Bill) IsDueBy(day int) bool {
	return b.DueDay <= day
}

// Label returns the bill as "Name: $D.CC", for example "Rent: $1200.50".
func (b Bill) Label() string {
	return fmt.Sprintf("%s: $%d.%02d", b.Name, b.AmountCents/100, b.AmountCents%100)
}

// TotalDueBy sums the amounts of all bills due on or before day.
func (bs Bills) TotalDueBy(day int) int64 {
	var total int64
	for _, b := range bs {
		if b.IsDueBy(day) {
			total += b.AmountCents
		}
	}
	return total
}
