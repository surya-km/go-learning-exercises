// Package errorwrapping practices the error wrapping concept. Read CONCEPT.md first.
package errorwrapping

// ProcessPayment validates the payment and charges the card. Errors are
// wrapped with the prefix "payment <ID>: ".
func ProcessPayment(p Payment) error {
	// TODO: implement
	return nil
}

// DeclineCode finds a *DeclinedError anywhere inside err and returns its Code.
func DeclineCode(err error) (string, bool) {
	// TODO: implement
	return "", false
}

// RootCause unwraps err layer by layer and returns the innermost error.
func RootCause(err error) error {
	// TODO: implement
	return nil
}

// ProcessBatch processes every payment and joins all failures into one error.
// It returns nil if nothing failed.
func ProcessBatch(payments []Payment) error {
	// TODO: implement
	return nil
}
