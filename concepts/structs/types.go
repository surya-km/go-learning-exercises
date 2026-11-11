package structs

// Account is a simple bank account. Money is stored in integer cents.
type Account struct {
	Owner        string
	Currency     string
	BalanceCents int64
}

// Transaction is one line on a bank statement.
// Positive amounts are money in, negative amounts are money out.
type Transaction struct {
	Description string
	AmountCents int64
}

// Summary is a short overview of a list of transactions.
type Summary struct {
	Count         int
	IncomeCents   int64
	SpendingCents int64
}
