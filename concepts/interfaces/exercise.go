// Package interfaces practices the interfaces concept. Read CONCEPT.md first.
package interfaces

// MonthlyCents returns the monthly price (yearly prices are divided by 12).
func (s Subscription) MonthlyCents() int64 {
	// TODO: implement
	return 0
}

// Name returns the service name.
func (s Subscription) Name() string {
	// TODO: implement
	return ""
}

// MonthlyCents returns the principal split evenly over the months (0 if Months <= 0).
func (l Loan) MonthlyCents() int64 {
	// TODO: implement
	return 0
}

// Name returns the lender followed by " loan".
func (l Loan) Name() string {
	// TODO: implement
	return ""
}

// MonthlyCents returns the average of the readings (0 if there are none).
func (u Utility) MonthlyCents() int64 {
	// TODO: implement
	return 0
}

// Name returns the provider name.
func (u Utility) Name() string {
	// TODO: implement
	return ""
}

// MonthlyTotal adds up the monthly cost of every payment.
func MonthlyTotal(payments []Payment) int64 {
	// TODO: implement
	return 0
}

// Biggest returns the name of the most expensive payment per month.
// Ties go to the earliest payment; an empty list returns "".
func Biggest(payments []Payment) string {
	// TODO: implement
	return ""
}
