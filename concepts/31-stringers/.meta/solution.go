// Package stringers practices the stringers concept. Read CONCEPT.md first.
package stringers

import (
	"fmt"
	"strings"
)

// String formats the amount as dollars, like "$19.99" or "-$4.50".
func (m Money) String() string {
	cents := int64(m)
	sign := ""
	if cents < 0 {
		sign = "-"
		cents = -cents
	}
	return fmt.Sprintf("%s$%d.%02d", sign, cents/100, cents%100)
}

// String formats the transaction as "DATE DESCRIPTION AMOUNT".
func (t Transaction) String() string {
	return fmt.Sprintf("%s %s %s", t.Date, t.Description, t.Amount)
}

// String formats the statement as a header line plus one line per transaction.
func (s Statement) String() string {
	noun := "transactions"
	if len(s.Transactions) == 1 {
		noun = "transaction"
	}

	var b strings.Builder
	fmt.Fprintf(&b, "Statement for %s (%d %s)\n", s.Owner, len(s.Transactions), noun)
	for _, t := range s.Transactions {
		fmt.Fprintf(&b, "%s\n", t)
	}
	return b.String()
}
