package slicesmaps

import (
	"maps"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name    string
		amounts []int
		want    int
	}{
		{name: "several amounts", amounts: []int{100, 250, 75}, want: 425},
		{name: "single amount", amounts: []int{42}, want: 42},
		{name: "includes refunds (negatives)", amounts: []int{500, -200}, want: 300},
		{name: "empty slice", amounts: []int{}, want: 0},
		{name: "nil slice", amounts: nil, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.amounts); got != tt.want {
				t.Errorf("Sum(%v) = %d, want %d", tt.amounts, got, tt.want)
			}
		})
	}
}

func TestAbove(t *testing.T) {
	tests := []struct {
		name    string
		amounts []int
		min     int
		want    []int
	}{
		{name: "keeps larger amounts in order", amounts: []int{50, 500, 20, 900}, min: 100, want: []int{500, 900}},
		{name: "equal to min is excluded", amounts: []int{100, 101}, min: 100, want: []int{101}},
		{name: "nothing matches", amounts: []int{1, 2, 3}, min: 10, want: []int{}},
		{name: "empty input", amounts: nil, min: 0, want: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Above(tt.amounts, tt.min)
			if len(got) != len(tt.want) || !slices.Equal(got, tt.want) {
				t.Errorf("Above(%v, %d) = %v, want %v", tt.amounts, tt.min, got, tt.want)
			}
		})
	}
}

func TestAboveDoesNotModifyInput(t *testing.T) {
	input := []int{5, 50, 500}
	original := slices.Clone(input)
	Above(input, 10)
	if !slices.Equal(input, original) {
		t.Errorf("Above modified its input: got %v, want %v", input, original)
	}
}

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
				t.Fatal("TotalByCategory returned a nil map; return an empty map instead")
			}
			if !maps.Equal(got, tt.want) {
				t.Errorf("TotalByCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}
