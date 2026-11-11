// Package constants practices the constants concept. Read CONCEPT.md first.
package constants

// Frequency says how often a repeating payment happens.
type Frequency int

// The possible frequencies. Unknown is the zero value.
const (
	Unknown Frequency = iota
	Weekly
	Fortnightly
	Monthly
	Yearly
)

// How many payments happen in one year at each frequency.
const (
	weeklyPaymentsPerYear      = 52
	fortnightlyPaymentsPerYear = 26
	monthlyPaymentsPerYear     = 12
	yearlyPaymentsPerYear      = 1
)

// YearlyCostCents returns how much a payment of amountCents made at
// frequency f costs over one year. Unknown frequencies cost 0.
func YearlyCostCents(amountCents int, f Frequency) int {
	switch f {
	case Weekly:
		return amountCents * weeklyPaymentsPerYear
	case Fortnightly:
		return amountCents * fortnightlyPaymentsPerYear
	case Monthly:
		return amountCents * monthlyPaymentsPerYear
	case Yearly:
		return amountCents * yearlyPaymentsPerYear
	default:
		return 0
	}
}
