// Package constants practices the constants concept. Read CONCEPT.md first.
package constants

// Frequency says how often a repeating payment happens.
type Frequency int

// The possible frequencies. Unknown is the zero value.
//
// TODO: these are all placeholders set to 0. Rewrite the block using iota.
const (
	Unknown     Frequency = 0
	Weekly      Frequency = 0
	Fortnightly Frequency = 0
	Monthly     Frequency = 0
	Yearly      Frequency = 0
)

// YearlyCostCents returns how much a payment of amountCents made at
// frequency f costs over one year. Unknown frequencies cost 0.
func YearlyCostCents(amountCents int, f Frequency) int {
	// TODO: implement
	return 0
}
