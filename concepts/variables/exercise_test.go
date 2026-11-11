package variables

import "testing"

func TestMonthEnd(t *testing.T) {
	tests := []struct {
		name           string
		startCents     int
		incomeCents    int
		rentCents      int
		groceriesCents int
		wantEnd        int
		wantSpent      int
	}{
		{name: "typical month", startCents: 50000, incomeCents: 300000, rentCents: 120000, groceriesCents: 40000, wantEnd: 190000, wantSpent: 160000},
		{name: "no income", startCents: 200000, incomeCents: 0, rentCents: 120000, groceriesCents: 30000, wantEnd: 50000, wantSpent: 150000},
		{name: "ends overdrawn", startCents: 0, incomeCents: 100000, rentCents: 120000, groceriesCents: 10000, wantEnd: -30000, wantSpent: 130000},
		{name: "spent nothing", startCents: 1000, incomeCents: 2000, rentCents: 0, groceriesCents: 0, wantEnd: 3000, wantSpent: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			end, spent := MonthEnd(tt.startCents, tt.incomeCents, tt.rentCents, tt.groceriesCents)
			if end != tt.wantEnd || spent != tt.wantSpent {
				t.Errorf("MonthEnd(%d, %d, %d, %d) = (%d, %d), want (%d, %d)",
					tt.startCents, tt.incomeCents, tt.rentCents, tt.groceriesCents,
					end, spent, tt.wantEnd, tt.wantSpent)
			}
		})
	}
}

func TestTransfer(t *testing.T) {
	tests := []struct {
		name        string
		fromCents   int
		toCents     int
		amountCents int
		wantFrom    int
		wantTo      int
	}{
		{name: "normal transfer", fromCents: 1000, toCents: 500, amountCents: 300, wantFrom: 700, wantTo: 800},
		{name: "move everything exactly", fromCents: 1000, toCents: 0, amountCents: 1000, wantFrom: 0, wantTo: 1000},
		{name: "not enough money", fromCents: 1000, toCents: 500, amountCents: 5000, wantFrom: 0, wantTo: 1500},
		{name: "zero amount", fromCents: 1000, toCents: 500, amountCents: 0, wantFrom: 1000, wantTo: 500},
		{name: "negative amount", fromCents: 1000, toCents: 500, amountCents: -200, wantFrom: 1000, wantTo: 500},
		{name: "empty source", fromCents: 0, toCents: 500, amountCents: 100, wantFrom: 0, wantTo: 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from, to := Transfer(tt.fromCents, tt.toCents, tt.amountCents)
			if from != tt.wantFrom || to != tt.wantTo {
				t.Errorf("Transfer(%d, %d, %d) = (%d, %d), want (%d, %d)",
					tt.fromCents, tt.toCents, tt.amountCents, from, to, tt.wantFrom, tt.wantTo)
			}
		})
	}
}
