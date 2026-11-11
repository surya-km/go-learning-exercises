// Package typeassertion practices the type assertion concept. Read CONCEPT.md first.
package typeassertion

import "fmt"

// Describe returns a human-readable description of a bank event.
// Use a type switch over the possible event types.
func Describe(event any) string {
	switch e := event.(type) {
	case Deposit:
		return fmt.Sprintf("deposit of %d cents from %s", e.AmountCents, e.From)
	case Withdrawal:
		if e.ATM {
			return fmt.Sprintf("ATM withdrawal of %d cents", e.AmountCents)
		}
		return fmt.Sprintf("withdrawal of %d cents", e.AmountCents)
	case Fee:
		return fmt.Sprintf("fee of %d cents (%s)", e.AmountCents, e.Reason)
	case string:
		return "note: " + e
	case nil:
		return "empty event"
	default:
		return fmt.Sprintf("unknown event of type %T", e)
	}
}

// LookupString returns payload[key] if it exists and is a string.
func LookupString(payload map[string]any, key string) (string, bool) {
	s, ok := payload[key].(string)
	return s, ok
}

// LookupCents returns payload[key] if it exists and is an int64.
func LookupCents(payload map[string]any, key string) (int64, bool) {
	n, ok := payload[key].(int64)
	return n, ok
}
