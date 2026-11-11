package randomness

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
	"testing"
)

// newRand returns a generator with a fixed seed so every test run is the same.
func newRand(seed uint64) *rand.Rand {
	return rand.New(rand.NewPCG(seed, 2))
}

func TestVerificationCodeShape(t *testing.T) {
	tests := []struct {
		length  int
		wantLen int
	}{
		{length: 1, wantLen: 1},
		{length: 4, wantLen: 4},
		{length: 6, wantLen: 6},
		{length: 32, wantLen: 32},
		{length: 0, wantLen: 0},
		{length: -3, wantLen: 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("length_%d", tt.length), func(t *testing.T) {
			r := newRand(1)
			for range 50 {
				got := VerificationCode(r, tt.length)
				if len(got) != tt.wantLen {
					t.Fatalf("VerificationCode(r, %d) = %q (length %d), want length %d", tt.length, got, len(got), tt.wantLen)
				}
				if strings.Trim(got, "0123456789") != "" {
					t.Fatalf("VerificationCode(r, %d) = %q, want only digits 0-9", tt.length, got)
				}
			}
		})
	}
}

func TestVerificationCodeUsesEveryDigit(t *testing.T) {
	r := newRand(7)
	seen := map[rune]bool{}
	for range 100 {
		for _, d := range VerificationCode(r, 6) {
			seen[d] = true
		}
	}
	for _, d := range "0123456789" {
		if !seen[d] {
			t.Errorf("digit %q never appeared in 100 codes of length 6; got digits %v", d, seen)
		}
	}
}

func TestVerificationCodeUsesGenerator(t *testing.T) {
	t.Run("same seed gives same codes", func(t *testing.T) {
		a, b := newRand(42), newRand(42)
		for i := range 5 {
			codeA, codeB := VerificationCode(a, 8), VerificationCode(b, 8)
			if codeA != codeB {
				t.Fatalf("call %d with seed 42: got %q and %q, want identical codes (use r, not the top-level rand functions)", i, codeA, codeB)
			}
		}
	})

	t.Run("codes vary between calls", func(t *testing.T) {
		r := newRand(3)
		first := VerificationCode(r, 8)
		for range 20 {
			if VerificationCode(r, 8) != first {
				return
			}
		}
		t.Errorf("VerificationCode(r, 8) returned %q 21 times in a row, want different codes", first)
	})
}

func TestShuffledCopyKeepsItems(t *testing.T) {
	tests := []struct {
		name  string
		items []string
	}{
		{name: "five transactions", items: []string{"txn1", "txn2", "txn3", "txn4", "txn5"}},
		{name: "with duplicates", items: []string{"rent", "food", "food", "fun"}},
		{name: "one item", items: []string{"only"}},
		{name: "empty", items: []string{}},
		{name: "nil", items: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := slices.Clone(tt.items)
			got := ShuffledCopy(newRand(1), tt.items)

			if !slices.Equal(tt.items, original) {
				t.Fatalf("ShuffledCopy changed its input: now %q, want %q", tt.items, original)
			}
			if len(got) != len(tt.items) {
				t.Fatalf("ShuffledCopy(r, %q) = %q (length %d), want length %d", tt.items, got, len(got), len(tt.items))
			}
			sortedGot, sortedWant := slices.Sorted(slices.Values(got)), slices.Sorted(slices.Values(tt.items))
			if !slices.Equal(sortedGot, sortedWant) {
				t.Errorf("ShuffledCopy(r, %q) = %q, want the same items in any order", tt.items, got)
			}
		})
	}
}

func TestShuffledCopyShuffles(t *testing.T) {
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	t.Run("same seed gives same order", func(t *testing.T) {
		got1 := ShuffledCopy(newRand(99), items)
		got2 := ShuffledCopy(newRand(99), items)
		if len(got1) != len(items) || !slices.Equal(got1, got2) {
			t.Fatalf("seed 99 gave %q and then %q, want two identical shuffles of %q", got1, got2, items)
		}
	})

	t.Run("order actually changes", func(t *testing.T) {
		r := newRand(5)
		orders := map[string]bool{}
		for range 20 {
			orders[strings.Join(ShuffledCopy(r, items), "")] = true
		}
		if len(orders) < 10 {
			t.Errorf("20 shuffles of %q produced only %d different orders %v, want at least 10", items, len(orders), orders)
		}
	})
}
