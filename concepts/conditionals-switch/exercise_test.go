package conditionalsswitch

import "testing"

func TestCategory(t *testing.T) {
	tests := []struct {
		merchant string
		want     string
	}{
		{merchant: "Tesco", want: "Groceries"},
		{merchant: "Aldi", want: "Groceries"},
		{merchant: "Lidl", want: "Groceries"},
		{merchant: "Shell", want: "Transport"},
		{merchant: "BP", want: "Transport"},
		{merchant: "Uber", want: "Transport"},
		{merchant: "Netflix", want: "Entertainment"},
		{merchant: "Spotify", want: "Entertainment"},
		{merchant: "Amazon", want: "Other"},
		{merchant: "tesco", want: "Other"},
		{merchant: "", want: "Other"},
	}

	for _, tt := range tests {
		t.Run("merchant="+tt.merchant, func(t *testing.T) {
			if got := Category(tt.merchant); got != tt.want {
				t.Errorf("Category(%q) = %q, want %q", tt.merchant, got, tt.want)
			}
		})
	}
}

func TestCreditRating(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  string
	}{
		{name: "top score", score: 850, want: "Excellent"},
		{name: "excellent lower edge", score: 800, want: "Excellent"},
		{name: "very good upper edge", score: 799, want: "Very good"},
		{name: "very good lower edge", score: 740, want: "Very good"},
		{name: "good upper edge", score: 739, want: "Good"},
		{name: "good lower edge", score: 670, want: "Good"},
		{name: "fair upper edge", score: 669, want: "Fair"},
		{name: "fair lower edge", score: 580, want: "Fair"},
		{name: "poor upper edge", score: 579, want: "Poor"},
		{name: "lowest valid score", score: 300, want: "Poor"},
		{name: "just below valid", score: 299, want: "Invalid"},
		{name: "just above valid", score: 851, want: "Invalid"},
		{name: "zero", score: 0, want: "Invalid"},
		{name: "negative", score: -10, want: "Invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreditRating(tt.score); got != tt.want {
				t.Errorf("CreditRating(%d) = %q, want %q", tt.score, got, tt.want)
			}
		})
	}
}
