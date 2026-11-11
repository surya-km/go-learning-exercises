package stringers

// Money is an amount in US cents.
type Money int64

// Transaction is one line on a bank statement.
type Transaction struct {
	Date        string // already formatted, like "2024-03-01"
	Description string
	Amount      Money
}

// Statement is a list of transactions for one account owner.
type Statement struct {
	Owner        string
	Transactions []Transaction
}
