// Package stringspackage practices the strings-package concept. Read CONCEPT.md first.
package stringspackage

// CleanMerchant tidies a raw bank statement merchant name: it drops a "POS " or "SQ *"
// prefix, cuts off store numbers starting at " #", collapses spaces and upper-cases it.
func CleanMerchant(raw string) string {
	// TODO: implement
	return ""
}

// ParseTags splits a comma-separated list into trimmed, lower-cased, unique, non-empty tags,
// in the order they first appear.
func ParseTags(s string) []string {
	// TODO: implement
	return nil
}
