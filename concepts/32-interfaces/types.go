package interfaces

// Payment is anything the budgeting app can count as a monthly cost.
type Payment interface {
	MonthlyCents() int64
	Name() string
}

// Subscription is a streaming service, gym, app, etc.
// If Yearly is true, PriceCents is the price for a whole year.
type Subscription struct {
	Service    string
	PriceCents int64
	Yearly     bool
}

// Loan is paid back in equal monthly instalments.
type Loan struct {
	Lender         string
	PrincipalCents int64
	Months         int64
}

// Utility is a bill that changes every month (electricity, water...).
// ReadingsCents holds the amounts of the last few bills.
type Utility struct {
	Provider      string
	ReadingsCents []int64
}
