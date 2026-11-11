// Package interfaces practices the interfaces concept. Read CONCEPT.md first.
package interfaces

// MonthlyCents returns the monthly price (yearly prices are divided by 12).
func (s Subscription) MonthlyCents() int64 {
	if s.Yearly {
		return s.PriceCents / 12
	}
	return s.PriceCents
}

// Name returns the service name.
func (s Subscription) Name() string {
	return s.Service
}

// MonthlyCents returns the principal split evenly over the months (0 if Months <= 0).
func (l Loan) MonthlyCents() int64 {
	if l.Months <= 0 {
		return 0
	}
	return l.PrincipalCents / l.Months
}

// Name returns the lender followed by " loan".
func (l Loan) Name() string {
	return l.Lender + " loan"
}

// MonthlyCents returns the average of the readings (0 if there are none).
func (u Utility) MonthlyCents() int64 {
	if len(u.ReadingsCents) == 0 {
		return 0
	}
	var sum int64
	for _, r := range u.ReadingsCents {
		sum += r
	}
	return sum / int64(len(u.ReadingsCents))
}

// Name returns the provider name.
func (u Utility) Name() string {
	return u.Provider
}

// MonthlyTotal adds up the monthly cost of every payment.
func MonthlyTotal(payments []Payment) int64 {
	var total int64
	for _, p := range payments {
		total += p.MonthlyCents()
	}
	return total
}

// Biggest returns the name of the most expensive payment per month.
// Ties go to the earliest payment; an empty list returns "".
func Biggest(payments []Payment) string {
	if len(payments) == 0 {
		return ""
	}
	best := payments[0]
	for _, p := range payments[1:] {
		if p.MonthlyCents() > best.MonthlyCents() {
			best = p
		}
	}
	return best.Name()
}
