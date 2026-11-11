package runes

import (
	"fmt"
	"testing"
)

func TestTruncate(t *testing.T) {
	tests := []struct {
		s    string
		max  int
		want string
	}{
		{s: "Coffee", max: 10, want: "Coffee"},
		{s: "Coffee", max: 6, want: "Coffee"},
		{s: "Starbucks Coffee", max: 10, want: "Starbucks…"},
		{s: "Café Crème", max: 6, want: "Café …"},
		{s: "Crème brûlée", max: 12, want: "Crème brûlée"},
		{s: "Crème brûlée", max: 9, want: "Crème br…"},
		{s: "😀😀😀", max: 2, want: "😀…"},
		{s: "€€€€", max: 3, want: "€€…"},
		{s: "abc", max: 1, want: "…"},
		{s: "abc", max: 0, want: ""},
		{s: "", max: 5, want: ""},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_max_%d", tt.s, tt.max), func(t *testing.T) {
			got := Truncate(tt.s, tt.max)
			if got != tt.want {
				t.Errorf("Truncate(%q, %d) = %q, want %q", tt.s, tt.max, got, tt.want)
			}
		})
	}
}

func TestInitials(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "ada lovelace", want: "AL"},
		{name: "Grace Brewster Hopper", want: "GBH"},
		{name: "Émile Zola", want: "ÉZ"},
		{name: "  élodie   dupont ", want: "ÉD"},
		{name: "ørjan", want: "Ø"},
		{name: "zoë\tñúñez", want: "ZÑ"},
		{name: "", want: ""},
		{name: "   ", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Initials(tt.name)
			if got != tt.want {
				t.Errorf("Initials(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
