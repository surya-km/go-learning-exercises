// Package typeassertion practices the type assertion concept. Read CONCEPT.md first.
package typeassertion

// Describe returns a human-readable description of a bank event.
// Use a type switch over the possible event types.
func Describe(event any) string {
	// TODO: implement
	return ""
}

// LookupString returns payload[key] if it exists and is a string.
func LookupString(payload map[string]any, key string) (string, bool) {
	// TODO: implement
	return "", false
}

// LookupCents returns payload[key] if it exists and is an int64.
func LookupCents(payload map[string]any, key string) (int64, bool) {
	// TODO: implement
	return 0, false
}
