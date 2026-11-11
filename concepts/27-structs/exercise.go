// Package structs practices the structs concept. Read CONCEPT.md first.
package structs

// NewAccount builds a USD account for owner (spaces trimmed).
// A negative openingCents is treated as 0.
func NewAccount(owner string, openingCents int64) Account {
	// TODO: implement
	return Account{}
}

// Summarize counts the transactions and totals money in and money out.
// SpendingCents is reported as a positive number.
func Summarize(txns []Transaction) Summary {
	// TODO: implement
	return Summary{}
}
