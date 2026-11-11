// Package methods practices the methods concept. Read CONCEPT.md first.
package methods

// IsDueBy reports whether the bill is due on or before day.
func (b Bill) IsDueBy(day int) bool {
	// TODO: implement
	return false
}

// Label returns the bill as "Name: $D.CC", for example "Rent: $1200.50".
func (b Bill) Label() string {
	// TODO: implement
	return ""
}

// TotalDueBy sums the amounts of all bills due on or before day.
func (bs Bills) TotalDueBy(day int) int64 {
	// TODO: implement
	return 0
}
