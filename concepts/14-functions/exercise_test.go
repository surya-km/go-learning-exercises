package functions

import "testing"

func TestGreeting(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "with a name", input: "Ada", want: "Hello, Ada!"},
		{name: "another name", input: "Grace", want: "Hello, Grace!"},
		{name: "name with a space", input: "Mary Ann", want: "Hello, Mary Ann!"},
		{name: "empty name falls back to friend", input: "", want: "Hello, friend!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greeting(tt.input); got != tt.want {
				t.Errorf("Greeting(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestTipCents(t *testing.T) {
	tests := []struct {
		name      string
		billCents int
		percent   int
		want      int
	}{
		{name: "15 percent", billCents: 2000, percent: 15, want: 300},
		{name: "20 percent", billCents: 4550, percent: 20, want: 910},
		{name: "rounds down", billCents: 1999, percent: 15, want: 299},
		{name: "zero percent", billCents: 2000, percent: 0, want: 0},
		{name: "negative percent means no tip", billCents: 2000, percent: -5, want: 0},
		{name: "zero bill", billCents: 0, percent: 18, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TipCents(tt.billCents, tt.percent); got != tt.want {
				t.Errorf("TipCents(%d, %d) = %d, want %d", tt.billCents, tt.percent, got, tt.want)
			}
		})
	}
}

func TestTotalWithTip(t *testing.T) {
	tests := []struct {
		name      string
		billCents int
		percent   int
		want      int
	}{
		{name: "20 percent", billCents: 2000, percent: 20, want: 2400},
		{name: "rounds the tip down", billCents: 1999, percent: 15, want: 2298},
		{name: "no tip", billCents: 2000, percent: 0, want: 2000},
		{name: "negative percent means no tip", billCents: 1500, percent: -10, want: 1500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalWithTip(tt.billCents, tt.percent); got != tt.want {
				t.Errorf("TotalWithTip(%d, %d) = %d, want %d", tt.billCents, tt.percent, got, tt.want)
			}
		})
	}
}
