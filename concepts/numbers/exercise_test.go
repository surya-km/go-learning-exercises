package numbers

import (
	"fmt"
	"slices"
	"testing"
)

func TestSplitBill(t *testing.T) {
	tests := []struct {
		total, people int
		want          []int
	}{
		{total: 1000, people: 4, want: []int{250, 250, 250, 250}},
		{total: 1000, people: 3, want: []int{334, 333, 333}},
		{total: 1001, people: 3, want: []int{334, 334, 333}},
		{total: 2, people: 4, want: []int{1, 1, 0, 0}},
		{total: 0, people: 2, want: []int{0, 0}},
		{total: 4599, people: 1, want: []int{4599}},
		{total: 500, people: 0, want: nil},
		{total: 500, people: -2, want: nil},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d_cents_%d_people", tt.total, tt.people), func(t *testing.T) {
			got := SplitBill(tt.total, tt.people)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SplitBill(%d, %d) = %v, want %v", tt.total, tt.people, got, tt.want)
			}
		})
	}
}

func TestMakeChange(t *testing.T) {
	tests := []struct {
		cents                             int
		quarters, dimes, nickels, pennies int
	}{
		{cents: 41, quarters: 1, dimes: 1, nickels: 1, pennies: 1},
		{cents: 99, quarters: 3, dimes: 2, nickels: 0, pennies: 4},
		{cents: 30, quarters: 1, dimes: 0, nickels: 1, pennies: 0},
		{cents: 100, quarters: 4, dimes: 0, nickels: 0, pennies: 0},
		{cents: 4, quarters: 0, dimes: 0, nickels: 0, pennies: 4},
		{cents: 0},
		{cents: -5},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d_cents", tt.cents), func(t *testing.T) {
			q, d, n, p := MakeChange(tt.cents)
			if q != tt.quarters || d != tt.dimes || n != tt.nickels || p != tt.pennies {
				t.Errorf("MakeChange(%d) = %d, %d, %d, %d; want %d, %d, %d, %d",
					tt.cents, q, d, n, p, tt.quarters, tt.dimes, tt.nickels, tt.pennies)
			}
		})
	}
}
