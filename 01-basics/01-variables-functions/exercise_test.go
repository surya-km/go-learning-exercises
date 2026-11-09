package variablesfunctions

import "testing"

func TestGreeting(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "with a name", input: "Ada", want: "Hello, Ada!"},
		{name: "another name", input: "Grace", want: "Hello, Grace!"},
		{name: "empty name falls back to friend", input: "", want: "Hello, friend!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Greeting(tt.input)
			if got != tt.want {
				t.Errorf("Greeting(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

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
