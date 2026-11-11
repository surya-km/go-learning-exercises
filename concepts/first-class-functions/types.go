package firstclassfunctions

// Transaction is a single spending record. Amounts are in cents.
type Transaction struct {
	Merchant    string
	AmountCents int
}
