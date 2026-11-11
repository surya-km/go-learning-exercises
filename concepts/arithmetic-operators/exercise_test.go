package arithmeticoperators

import "testing"

func TestTipCents(t *testing.T) {
	tests := []struct {
		name      string
		billCents int
		percent   int
		want      int
	}{
		{name: "exact 15 percent", billCents: 1000, percent: 15, want: 150},
		{name: "rounds 399.8 up to 400", billCents: 1999, percent: 20, want: 400},
		{name: "rounds 222.12 down to 222", billCents: 1234, percent: 18, want: 222},
		{name: "half a cent rounds up", billCents: 50, percent: 1, want: 1},
		{name: "just under half rounds down", billCents: 49, percent: 1, want: 0},
		{name: "zero percent", billCents: 5000, percent: 0, want: 0},
		{name: "zero bill", billCents: 0, percent: 20, want: 0},
		{name: "100 percent", billCents: 4321, percent: 100, want: 4321},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TipCents(tt.billCents, tt.percent); got != tt.want {
				t.Errorf("TipCents(%d, %d) = %d, want %d", tt.billCents, tt.percent, got, tt.want)
			}
		})
	}
}

func TestRoundUp(t *testing.T) {
	tests := []struct {
		name        string
		cents       int
		wantRounded int
		wantSpare   int
	}{
		{name: "typical purchase", cents: 1234, wantRounded: 1300, wantSpare: 66},
		{name: "already whole dollar", cents: 1300, wantRounded: 1300, wantSpare: 0},
		{name: "one cent", cents: 1, wantRounded: 100, wantSpare: 99},
		{name: "99 cents", cents: 99, wantRounded: 100, wantSpare: 1},
		{name: "one cent over a dollar", cents: 101, wantRounded: 200, wantSpare: 99},
		{name: "zero", cents: 0, wantRounded: 0, wantSpare: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rounded, spare := RoundUp(tt.cents)
			if rounded != tt.wantRounded || spare != tt.wantSpare {
				t.Errorf("RoundUp(%d) = (%d, %d), want (%d, %d)",
					tt.cents, rounded, spare, tt.wantRounded, tt.wantSpare)
			}
		})
	}
}
