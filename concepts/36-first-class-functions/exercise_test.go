package firstclassfunctions

import (
	"slices"
	"strings"
	"testing"
)

var sampleTxns = []Transaction{
	{Merchant: "Grocer", AmountCents: 4520},
	{Merchant: "Landlord", AmountCents: 120000},
	{Merchant: "Coffee Shop", AmountCents: 450},
	{Merchant: "Gas Station", AmountCents: 6000},
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name string
		txns []Transaction
		desc string
		keep func(Transaction) bool
		want []Transaction
	}{
		{
			name: "amounts over $50",
			txns: sampleTxns,
			desc: "AmountCents > 5000",
			keep: func(t Transaction) bool { return t.AmountCents > 5000 },
			want: []Transaction{
				{Merchant: "Landlord", AmountCents: 120000},
				{Merchant: "Gas Station", AmountCents: 6000},
			},
		},
		{
			name: "merchant contains Shop",
			txns: sampleTxns,
			desc: `Merchant contains "Shop"`,
			keep: func(t Transaction) bool { return strings.Contains(t.Merchant, "Shop") },
			want: []Transaction{{Merchant: "Coffee Shop", AmountCents: 450}},
		},
		{
			name: "keep everything",
			txns: sampleTxns,
			desc: "always true",
			keep: func(Transaction) bool { return true },
			want: sampleTxns,
		},
		{
			name: "nothing matches",
			txns: sampleTxns,
			desc: "always false",
			keep: func(Transaction) bool { return false },
			want: []Transaction{},
		},
		{
			name: "no transactions",
			txns: nil,
			desc: "always true",
			keep: func(Transaction) bool { return true },
			want: []Transaction{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.txns, tt.keep)
			if got == nil {
				t.Fatalf("Filter(%v, %s) returned nil; return an empty slice instead", tt.txns, tt.desc)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("Filter(%v, %s) = %v, want %v", tt.txns, tt.desc, got, tt.want)
			}
		})
	}
}

func TestFilterDoesNotModifyInput(t *testing.T) {
	input := slices.Clone(sampleTxns)
	Filter(input, func(t Transaction) bool { return t.AmountCents < 1000 })
	if !slices.Equal(input, sampleTxns) {
		t.Errorf("Filter modified its input: got %v, want %v", input, sampleTxns)
	}
}

func TestSpendingTracker(t *testing.T) {
	type call struct {
		amount        int
		wantRemaining int
		wantOK        bool
	}
	tests := []struct {
		name  string
		limit int
		calls []call
	}{
		{
			name:  "spends until the budget runs out",
			limit: 5000,
			calls: []call{
				{amount: 2000, wantRemaining: 3000, wantOK: true},
				{amount: 4000, wantRemaining: 3000, wantOK: false},
				{amount: 3000, wantRemaining: 0, wantOK: true},
				{amount: 1, wantRemaining: 0, wantOK: false},
			},
		},
		{
			name:  "zero amount is always fine",
			limit: 100,
			calls: []call{
				{amount: 0, wantRemaining: 100, wantOK: true},
			},
		},
		{
			name:  "first purchase too big",
			limit: 1000,
			calls: []call{
				{amount: 1500, wantRemaining: 1000, wantOK: false},
				{amount: 999, wantRemaining: 1, wantOK: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spend := SpendingTracker(tt.limit)
			if spend == nil {
				t.Fatalf("SpendingTracker(%d) returned nil, want a function", tt.limit)
			}
			for i, c := range tt.calls {
				remaining, ok := spend(c.amount)
				if remaining != c.wantRemaining || ok != c.wantOK {
					t.Errorf("SpendingTracker(%d): call #%d spend(%d) = (%d, %t), want (%d, %t)",
						tt.limit, i+1, c.amount, remaining, ok, c.wantRemaining, c.wantOK)
				}
			}
		})
	}
}

func TestSpendingTrackersAreIndependent(t *testing.T) {
	alice := SpendingTracker(1000)
	bob := SpendingTracker(1000)
	if alice == nil || bob == nil {
		t.Fatal("SpendingTracker returned nil, want a function")
	}
	alice(800)
	if remaining, ok := bob(500); remaining != 500 || !ok {
		t.Errorf("after alice(800), bob(500) = (%d, %t), want (500, true); each tracker needs its own budget", remaining, ok)
	}
	if remaining, ok := alice(500); remaining != 200 || ok {
		t.Errorf("alice(500) = (%d, %t), want (200, false)", remaining, ok)
	}
}
