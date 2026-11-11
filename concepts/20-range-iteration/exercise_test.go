package rangeiteration

import (
	"slices"
	"testing"
)

func TestRunningBalance(t *testing.T) {
	tests := []struct {
		name       string
		startCents int
		amounts    []int
		want       []int
	}{
		{name: "deposits and withdrawals", startCents: 1000, amounts: []int{-200, 500, -50}, want: []int{800, 1300, 1250}},
		{name: "starts at zero", startCents: 0, amounts: []int{100, 100, 100}, want: []int{100, 200, 300}},
		{name: "goes negative", startCents: 100, amounts: []int{-300}, want: []int{-200}},
		{name: "no amounts", startCents: 500, amounts: nil, want: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RunningBalance(tt.startCents, tt.amounts)
			if got == nil || !slices.Equal(got, tt.want) {
				t.Errorf("RunningBalance(%d, %v) = %#v, want %v", tt.startCents, tt.amounts, got, tt.want)
			}
		})
	}
}

func TestInstallmentPlan(t *testing.T) {
	tests := []struct {
		name       string
		totalCents int
		months     int
		want       []int
	}{
		{name: "splits evenly", totalCents: 1200, months: 4, want: []int{300, 300, 300, 300}},
		{name: "one leftover cent", totalCents: 1000, months: 3, want: []int{334, 333, 333}},
		{name: "two leftover cents", totalCents: 1001, months: 3, want: []int{334, 334, 333}},
		{name: "single month", totalCents: 4999, months: 1, want: []int{4999}},
		{name: "more months than cents", totalCents: 2, months: 4, want: []int{1, 1, 0, 0}},
		{name: "zero months", totalCents: 1000, months: 0, want: []int{}},
		{name: "negative months", totalCents: 1000, months: -1, want: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InstallmentPlan(tt.totalCents, tt.months)
			if got == nil || !slices.Equal(got, tt.want) {
				t.Errorf("InstallmentPlan(%d, %d) = %#v, want %v", tt.totalCents, tt.months, got, tt.want)
			}
		})
	}
}

func TestMaskCard(t *testing.T) {
	tests := []struct {
		name string
		card string
		want string
	}{
		{name: "spaced card number", card: "4111 1111 1111 1234", want: "**** **** **** 1234"},
		{name: "no separators", card: "5500000000000004", want: "************0004"},
		{name: "dashes", card: "3782-822463-10005", want: "****-******-*0005"},
		{name: "exactly four digits", card: "1234", want: "1234"},
		{name: "fewer than four digits", card: "12", want: "12"},
		{name: "non-ASCII separator is kept", card: "1234•5678", want: "****•5678"},
		{name: "empty", card: "", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaskCard(tt.card); got != tt.want {
				t.Errorf("MaskCard(%q) = %q, want %q", tt.card, got, tt.want)
			}
		})
	}
}

func TestBiggestCategory(t *testing.T) {
	tests := []struct {
		name      string
		totals    map[string]int
		wantName  string
		wantTotal int
	}{
		{name: "clear winner", totals: map[string]int{"food": 40000, "rent": 120000, "fun": 9000}, wantName: "rent", wantTotal: 120000},
		{name: "single category", totals: map[string]int{"travel": 300}, wantName: "travel", wantTotal: 300},
		{name: "tie goes to alphabetical first", totals: map[string]int{"rent": 500, "food": 500, "zoo": 500, "bills": 100}, wantName: "food", wantTotal: 500},
		{name: "all negative (refunds)", totals: map[string]int{"food": -500, "fun": -100}, wantName: "fun", wantTotal: -100},
		{name: "empty map", totals: map[string]int{}, wantName: "", wantTotal: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run several times: map order is random, so a lucky order shouldn't hide a bug.
			for range 20 {
				gotName, gotTotal := BiggestCategory(tt.totals)
				if gotName != tt.wantName || gotTotal != tt.wantTotal {
					t.Fatalf("BiggestCategory(%v) = (%q, %d), want (%q, %d)",
						tt.totals, gotName, gotTotal, tt.wantName, tt.wantTotal)
				}
			}
		})
	}
}
