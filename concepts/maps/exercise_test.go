package mapspractice

import (
	"maps"
	"slices"
	"testing"
)

func TestTotalByCategory(t *testing.T) {
	tests := []struct {
		name string
		txns []Transaction
		want map[string]int
	}{
		{
			name: "groups and sums by category",
			txns: []Transaction{
				{Category: "food", AmountCents: 1250},
				{Category: "rent", AmountCents: 120000},
				{Category: "food", AmountCents: 750},
			},
			want: map[string]int{"food": 2000, "rent": 120000},
		},
		{
			name: "refund reduces the total",
			txns: []Transaction{
				{Category: "shopping", AmountCents: 5000},
				{Category: "shopping", AmountCents: -1500},
			},
			want: map[string]int{"shopping": 3500},
		},
		{
			name: "single transaction",
			txns: []Transaction{{Category: "travel", AmountCents: 300}},
			want: map[string]int{"travel": 300},
		},
		{
			name: "no transactions",
			txns: nil,
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TotalByCategory(tt.txns)
			if got == nil {
				t.Fatalf("TotalByCategory(%v) returned a nil map; return an empty map instead", tt.txns)
			}
			if !maps.Equal(got, tt.want) {
				t.Errorf("TotalByCategory(%v) = %v, want %v", tt.txns, got, tt.want)
			}
		})
	}
}

func TestOverBudget(t *testing.T) {
	tests := []struct {
		name   string
		spent  map[string]int
		budget map[string]int
		want   []string
	}{
		{
			name:   "over budget and unbudgeted, sorted",
			spent:  map[string]int{"food": 45000, "rent": 120000, "games": 2000},
			budget: map[string]int{"food": 40000, "rent": 120000},
			want:   []string{"food", "games"},
		},
		{
			name:   "many over, sorted alphabetically",
			spent:  map[string]int{"zoo": 10, "fun": 10, "bills": 10, "auto": 10},
			budget: map[string]int{"zoo": 5, "fun": 5, "bills": 5, "auto": 5},
			want:   []string{"auto", "bills", "fun", "zoo"},
		},
		{
			name:   "exactly on budget is fine",
			spent:  map[string]int{"food": 40000},
			budget: map[string]int{"food": 40000},
			want:   []string{},
		},
		{
			name:   "unbudgeted with zero spent is fine",
			spent:  map[string]int{"gifts": 0},
			budget: map[string]int{},
			want:   []string{},
		},
		{
			name:   "zero budget, something spent",
			spent:  map[string]int{"gifts": 1},
			budget: map[string]int{"gifts": 0},
			want:   []string{"gifts"},
		},
		{
			name:   "budget-only categories are ignored",
			spent:  map[string]int{},
			budget: map[string]int{"travel": 5000},
			want:   []string{},
		},
		{
			name:   "nil maps",
			spent:  nil,
			budget: nil,
			want:   []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OverBudget(tt.spent, tt.budget)
			if got == nil || !slices.Equal(got, tt.want) {
				t.Errorf("OverBudget(%v, %v) = %#v, want %q", tt.spent, tt.budget, got, tt.want)
			}
		})
	}
}
