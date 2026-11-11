package constants

import "testing"

func TestFrequencyValues(t *testing.T) {
	tests := []struct {
		name string
		got  Frequency
		want Frequency
	}{
		{name: "Unknown", got: Unknown, want: 0},
		{name: "Weekly", got: Weekly, want: 1},
		{name: "Fortnightly", got: Fortnightly, want: 2},
		{name: "Monthly", got: Monthly, want: 3},
		{name: "Yearly", got: Yearly, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("%s = %d, want %d", tt.name, tt.got, tt.want)
			}
		})
	}
}

func TestYearlyCostCents(t *testing.T) {
	tests := []struct {
		name        string
		amountCents int
		f           Frequency
		want        int
	}{
		{name: "weekly coffee", amountCents: 350, f: Weekly, want: 18200},
		{name: "fortnightly cleaner", amountCents: 6000, f: Fortnightly, want: 156000},
		{name: "monthly streaming", amountCents: 1099, f: Monthly, want: 13188},
		{name: "yearly insurance", amountCents: 45000, f: Yearly, want: 45000},
		{name: "unknown frequency", amountCents: 1000, f: Unknown, want: 0},
		{name: "out of range frequency", amountCents: 1000, f: Frequency(99), want: 0},
		{name: "free monthly plan", amountCents: 0, f: Monthly, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YearlyCostCents(tt.amountCents, tt.f); got != tt.want {
				t.Errorf("YearlyCostCents(%d, Frequency(%d)) = %d, want %d",
					tt.amountCents, tt.f, got, tt.want)
			}
		})
	}
}
