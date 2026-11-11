package forloops

import "testing"

func TestMonthsToPayOff(t *testing.T) {
	tests := []struct {
		name               string
		balanceCents       int
		monthlyRatePercent int
		paymentCents       int
		want               int
	}{
		{name: "no interest", balanceCents: 1000, monthlyRatePercent: 0, paymentCents: 300, want: 4},
		{name: "exact payoff", balanceCents: 900, monthlyRatePercent: 0, paymentCents: 300, want: 3},
		{name: "one big payment", balanceCents: 500, monthlyRatePercent: 2, paymentCents: 10000, want: 1},
		{name: "with interest", balanceCents: 10000, monthlyRatePercent: 1, paymentCents: 2500, want: 5},
		{name: "high interest", balanceCents: 100000, monthlyRatePercent: 3, paymentCents: 10000, want: 13},
		{name: "payment equals interest", balanceCents: 10000, monthlyRatePercent: 5, paymentCents: 500, want: -1},
		{name: "payment below interest", balanceCents: 10000, monthlyRatePercent: 10, paymentCents: 100, want: -1},
		{name: "zero payment", balanceCents: 1000, monthlyRatePercent: 0, paymentCents: 0, want: -1},
		{name: "nothing owed", balanceCents: 0, monthlyRatePercent: 3, paymentCents: 100, want: 0},
		{name: "in credit", balanceCents: -50, monthlyRatePercent: 3, paymentCents: 100, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MonthsToPayOff(tt.balanceCents, tt.monthlyRatePercent, tt.paymentCents)
			if got != tt.want {
				t.Errorf("MonthsToPayOff(%d, %d, %d) = %d, want %d",
					tt.balanceCents, tt.monthlyRatePercent, tt.paymentCents, got, tt.want)
			}
		})
	}
}

func TestCountWeekdays(t *testing.T) {
	tests := []struct {
		name     string
		startDay int
		days     int
		want     int
	}{
		{name: "full week from Monday", startDay: 1, days: 7, want: 5},
		{name: "full week from Sunday", startDay: 0, days: 7, want: 5},
		{name: "weekend only", startDay: 6, days: 2, want: 0},
		{name: "Friday to Sunday", startDay: 5, days: 3, want: 1},
		{name: "ten days from Wednesday", startDay: 3, days: 10, want: 8},
		{name: "two weeks", startDay: 0, days: 14, want: 10},
		{name: "single weekday", startDay: 2, days: 1, want: 1},
		{name: "zero days", startDay: 1, days: 0, want: 0},
		{name: "negative days", startDay: 1, days: -3, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountWeekdays(tt.startDay, tt.days); got != tt.want {
				t.Errorf("CountWeekdays(%d, %d) = %d, want %d", tt.startDay, tt.days, got, tt.want)
			}
		})
	}
}
