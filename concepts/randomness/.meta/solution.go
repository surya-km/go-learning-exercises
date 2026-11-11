// Package randomness practices the randomness concept. Read CONCEPT.md first.
package randomness

import (
	"math/rand/v2"
	"slices"
	"strings"
)

// VerificationCode returns length random digits generated with r, like "048213".
// It returns "" if length <= 0.
func VerificationCode(r *rand.Rand, length int) string {
	if length <= 0 {
		return ""
	}
	var sb strings.Builder
	for range length {
		sb.WriteByte(byte('0' + r.IntN(10)))
	}
	return sb.String()
}

// ShuffledCopy returns a new slice holding items in a random order chosen with r.
// It does not modify items.
func ShuffledCopy(r *rand.Rand, items []string) []string {
	shuffled := slices.Clone(items)
	r.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}
