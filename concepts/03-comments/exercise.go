// Package comments practices the comments concept. Read CONCEPT.md first.
package comments

// TODO: every exported name below needs a proper doc comment (see CONCEPT.md).

const CurrencySymbol = "$"

// represents one payment in or out of an account.
type Transaction struct {
	Description string
	AmountCents int
}

// formats cents
func FormatAmount(cents int) string {
	// TODO: implement
	return ""
}

func Describe(t Transaction) string {
	// TODO: implement
	return ""
}
