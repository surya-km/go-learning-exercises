// Package stringspackage practices the strings-package concept. Read CONCEPT.md first.
package stringspackage

import (
	"slices"
	"strings"
)

// CleanMerchant tidies a raw bank statement merchant name: it drops a "POS " or "SQ *"
// prefix, cuts off store numbers starting at " #", collapses spaces and upper-cases it.
func CleanMerchant(raw string) string {
	s := strings.TrimSpace(raw)

	for _, prefix := range []string{"POS ", "SQ *"} {
		if strings.HasPrefix(s, prefix) {
			s = strings.TrimPrefix(s, prefix)
			break
		}
	}

	s, _, _ = strings.Cut(s, " #")

	s = strings.Join(strings.Fields(s), " ")
	return strings.ToUpper(s)
}

// ParseTags splits a comma-separated list into trimmed, lower-cased, unique, non-empty tags,
// in the order they first appear.
func ParseTags(s string) []string {
	var tags []string
	for _, part := range strings.Split(s, ",") {
		tag := strings.ToLower(strings.TrimSpace(part))
		if tag == "" || slices.Contains(tags, tag) {
			continue
		}
		tags = append(tags, tag)
	}
	return tags
}
