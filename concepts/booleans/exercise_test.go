package booleans

import "testing"

func TestCanWithdraw(t *testing.T) {
	tests := []struct {
		name         string
		balanceCents int
		amountCents  int
		frozen       bool
		want         bool
	}{
		{name: "enough money", balanceCents: 5000, amountCents: 2000, frozen: false, want: true},
		{name: "whole balance", balanceCents: 5000, amountCents: 5000, frozen: false, want: true},
		{name: "one cent too much", balanceCents: 5000, amountCents: 5001, frozen: false, want: false},
		{name: "frozen account", balanceCents: 5000, amountCents: 2000, frozen: true, want: false},
		{name: "zero amount", balanceCents: 5000, amountCents: 0, frozen: false, want: false},
		{name: "negative amount", balanceCents: 5000, amountCents: -100, frozen: false, want: false},
		{name: "empty account", balanceCents: 0, amountCents: 1, frozen: false, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanWithdraw(tt.balanceCents, tt.amountCents, tt.frozen)
			if got != tt.want {
				t.Errorf("CanWithdraw(%d, %d, %t) = %t, want %t",
					tt.balanceCents, tt.amountCents, tt.frozen, got, tt.want)
			}
		})
	}
}

func TestNeedsReview(t *testing.T) {
	tests := []struct {
		name          string
		amountCents   int
		international bool
		newPayee      bool
		want          bool
	}{
		{name: "small domestic known payee", amountCents: 2500, international: false, newPayee: false, want: false},
		{name: "exactly the limit", amountCents: 100000, international: false, newPayee: false, want: true},
		{name: "just under the limit", amountCents: 99999, international: false, newPayee: false, want: false},
		{name: "large amount", amountCents: 250000, international: false, newPayee: false, want: true},
		{name: "international new payee", amountCents: 500, international: true, newPayee: true, want: true},
		{name: "international known payee", amountCents: 500, international: true, newPayee: false, want: false},
		{name: "domestic new payee", amountCents: 500, international: false, newPayee: true, want: false},
		{name: "large international new payee", amountCents: 500000, international: true, newPayee: true, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NeedsReview(tt.amountCents, tt.international, tt.newPayee)
			if got != tt.want {
				t.Errorf("NeedsReview(%d, %t, %t) = %t, want %t",
					tt.amountCents, tt.international, tt.newPayee, got, tt.want)
			}
		})
	}
}
