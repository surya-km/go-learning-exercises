package mapspractice

// Transaction is a single spending record. Amounts are in cents.
type Transaction struct {
	Category    string
	AmountCents int
}
