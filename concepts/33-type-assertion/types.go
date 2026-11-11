package typeassertion

// Deposit is money coming into the account.
type Deposit struct {
	From        string
	AmountCents int64
}

// Withdrawal is money taken out, either at an ATM or by other means.
type Withdrawal struct {
	ATM         bool
	AmountCents int64
}

// Fee is a charge from the bank.
type Fee struct {
	Reason      string
	AmountCents int64
}
