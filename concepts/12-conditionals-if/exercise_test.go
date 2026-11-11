package conditionalsif

import "testing"

func TestOverdraftFee(t *testing.T) {
	tests := []struct {
		name         string
		balanceCents int
		want         int
	}{
		{name: "in credit", balanceCents: 12000, want: 0},
		{name: "exactly zero", balanceCents: 0, want: 0},
		{name: "one cent overdrawn", balanceCents: -1, want: 0},
		{name: "exactly at grace limit", balanceCents: -500, want: 0},
		{name: "one cent past grace limit", balanceCents: -501, want: 2500},
		{name: "deeply overdrawn", balanceCents: -100000, want: 2500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OverdraftFee(tt.balanceCents); got != tt.want {
				t.Errorf("OverdraftFee(%d) = %d, want %d", tt.balanceCents, got, tt.want)
			}
		})
	}
}

func TestBudgetStatus(t *testing.T) {
	tests := []struct {
		name        string
		spentCents  int
		budgetCents int
		want        string
	}{
		{name: "well under budget", spentCents: 1000, budgetCents: 10000, want: "ok"},
		{name: "nothing spent", spentCents: 0, budgetCents: 10000, want: "ok"},
		{name: "just under 80 percent", spentCents: 7999, budgetCents: 10000, want: "ok"},
		{name: "exactly 80 percent", spentCents: 8000, budgetCents: 10000, want: "warning"},
		{name: "exactly on budget", spentCents: 10000, budgetCents: 10000, want: "warning"},
		{name: "one cent over", spentCents: 10001, budgetCents: 10000, want: "over"},
		{name: "zero budget", spentCents: 500, budgetCents: 0, want: "no budget"},
		{name: "negative budget", spentCents: 0, budgetCents: -100, want: "no budget"},
		{name: "80 percent of an odd budget", spentCents: 8, budgetCents: 10, want: "warning"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BudgetStatus(tt.spentCents, tt.budgetCents); got != tt.want {
				t.Errorf("BudgetStatus(%d, %d) = %q, want %q", tt.spentCents, tt.budgetCents, got, tt.want)
			}
		})
	}
}
