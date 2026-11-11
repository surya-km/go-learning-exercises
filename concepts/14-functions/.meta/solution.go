// Package functions practices the functions concept. Read CONCEPT.md first.
package functions

// Greeting returns "Hello, <name>!", or "Hello, friend!" when name is empty.
func Greeting(name string) string {
	if name == "" {
		name = "friend"
	}
	return "Hello, " + name + "!"
}

// TipCents returns the tip on billCents at the given percent, rounded down.
// A percent of 0 or less means no tip.
func TipCents(billCents, percent int) int {
	if percent <= 0 {
		return 0
	}
	return billCents * percent / 100
}

// TotalWithTip returns billCents plus the tip. It should call TipCents.
func TotalWithTip(billCents, percent int) int {
	return billCents + TipCents(billCents, percent)
}
