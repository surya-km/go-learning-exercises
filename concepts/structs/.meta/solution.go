// Package structs practices the structs concept. Read CONCEPT.md first.
package structs

import "strings"

// NewAccount builds a USD account for owner (spaces trimmed).
// A negative openingCents is treated as 0.
func NewAccount(owner string, openingCents int64) Account {
	if openingCents < 0 {
		openingCents = 0
	}
	return Account{
		Owner:        strings.TrimSpace(owner),
		Currency:     "USD",
		BalanceCents: openingCents,
	}
}

// Summarize counts the transactions and totals money in and money out.
// SpendingCents is reported as a positive number.
func Summarize(txns []Transaction) Summary {
	var s Summary
	for _, t := range txns {
		s.Count++
		switch {
		case t.AmountCents > 0:
			s.IncomeCents += t.AmountCents
		case t.AmountCents < 0:
			s.SpendingCents -= t.AmountCents
		}
	}
	return s
}
