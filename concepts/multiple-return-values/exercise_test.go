package multiplereturnvalues

import "testing"

func TestSplitBill(t *testing.T) {
	tests := []struct {
		name         string
		totalCents   int
		people       int
		wantEach     int
		wantLeftover int
	}{
		{name: "splits evenly", totalCents: 1000, people: 4, wantEach: 250, wantLeftover: 0},
		{name: "has leftover cents", totalCents: 1000, people: 3, wantEach: 333, wantLeftover: 1},
		{name: "one person pays everything", totalCents: 4599, people: 1, wantEach: 4599, wantLeftover: 0},
		{name: "more people than cents", totalCents: 2, people: 5, wantEach: 0, wantLeftover: 2},
		{name: "zero people", totalCents: 500, people: 0, wantEach: 0, wantLeftover: 500},
		{name: "negative people", totalCents: 500, people: -2, wantEach: 0, wantLeftover: 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			each, leftover := SplitBill(tt.totalCents, tt.people)
			if each != tt.wantEach || leftover != tt.wantLeftover {
				t.Errorf("SplitBill(%d, %d) = (%d, %d), want (%d, %d)",
					tt.totalCents, tt.people, each, leftover, tt.wantEach, tt.wantLeftover)
			}
		})
	}
}

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name                              string
		cents                             int
		quarters, dimes, nickels, pennies int
	}{
		{name: "one of each coin", cents: 41, quarters: 1, dimes: 1, nickels: 1, pennies: 1},
		{name: "99 cents", cents: 99, quarters: 3, dimes: 2, nickels: 0, pennies: 4},
		{name: "quarter and nickel", cents: 30, quarters: 1, dimes: 0, nickels: 1, pennies: 0},
		{name: "only pennies", cents: 4, quarters: 0, dimes: 0, nickels: 0, pennies: 4},
		{name: "exact dollar", cents: 100, quarters: 4, dimes: 0, nickels: 0, pennies: 0},
		{name: "zero", cents: 0, quarters: 0, dimes: 0, nickels: 0, pennies: 0},
		{name: "negative", cents: -5, quarters: 0, dimes: 0, nickels: 0, pennies: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q, d, n, p := CoinChange(tt.cents)
			if q != tt.quarters || d != tt.dimes || n != tt.nickels || p != tt.pennies {
				t.Errorf("CoinChange(%d) = (%d, %d, %d, %d), want (%d, %d, %d, %d)",
					tt.cents, q, d, n, p, tt.quarters, tt.dimes, tt.nickels, tt.pennies)
			}
		})
	}
}
