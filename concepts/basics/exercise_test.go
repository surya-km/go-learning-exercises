package basics

import "testing"

func TestGreeting(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "simple name", input: "Ada", want: "Hello, Ada!"},
		{name: "full name", input: "Grace Hopper", want: "Hello, Grace Hopper!"},
		{name: "empty name", input: "", want: "Hello, friend!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greeting(tt.input); got != tt.want {
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
		{name: "even split", totalCents: 900, people: 3, wantEach: 300, wantLeftover: 0},
		{name: "uneven split", totalCents: 1000, people: 3, wantEach: 333, wantLeftover: 1},
		{name: "one person", totalCents: 1234, people: 1, wantEach: 1234, wantLeftover: 0},
		{name: "more people than cents", totalCents: 5, people: 10, wantEach: 0, wantLeftover: 5},
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
