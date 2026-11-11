// Package randomness practices the randomness concept. Read CONCEPT.md first.
package randomness

import "math/rand/v2"

// VerificationCode returns length random digits generated with r, like "048213".
// It returns "" if length <= 0.
func VerificationCode(r *rand.Rand, length int) string {
	// TODO: implement
	return ""
}

// ShuffledCopy returns a new slice holding items in a random order chosen with r.
// It does not modify items.
func ShuffledCopy(r *rand.Rand, items []string) []string {
	// TODO: implement
	return nil
}
