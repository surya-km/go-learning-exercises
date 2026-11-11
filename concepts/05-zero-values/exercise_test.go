package zerovalues

import (
	"fmt"
	"testing"
)

// describe prints a slice so that nil and empty slices look different in messages.
func describe(s []string) string {
	if s == nil {
		return "nil"
	}
	return fmt.Sprintf("%q", s)
}

func sameStrings(a, b []string) bool {
	if (a == nil) != (b == nil) || len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestApplyDefaults(t *testing.T) {
	tests := []struct {
		name           string
		currency       string
		limit          int
		categories     []string
		wantCurrency   string
		wantLimit      int
		wantCategories []string
	}{
		{name: "everything unset", currency: "", limit: 0, categories: nil,
			wantCurrency: "USD", wantLimit: 50000, wantCategories: []string{"General"}},
		{name: "everything set", currency: "GBP", limit: 20000, categories: []string{"Food", "Rent"},
			wantCurrency: "GBP", wantLimit: 20000, wantCategories: []string{"Food", "Rent"}},
		{name: "only currency unset", currency: "", limit: 10000, categories: []string{"Bills"},
			wantCurrency: "USD", wantLimit: 10000, wantCategories: []string{"Bills"}},
		{name: "only limit unset", currency: "EUR", limit: 0, categories: []string{"Bills"},
			wantCurrency: "EUR", wantLimit: 50000, wantCategories: []string{"Bills"}},
		{name: "empty non-nil categories are kept", currency: "EUR", limit: 100, categories: []string{},
			wantCurrency: "EUR", wantLimit: 100, wantCategories: []string{}},
		{name: "negative limit is a real value", currency: "USD", limit: -1, categories: nil,
			wantCurrency: "USD", wantLimit: -1, wantCategories: []string{"General"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cur, limit, cats := ApplyDefaults(tt.currency, tt.limit, tt.categories)
			if cur != tt.wantCurrency || limit != tt.wantLimit || !sameStrings(cats, tt.wantCategories) {
				t.Errorf("ApplyDefaults(%q, %d, %s) = (%q, %d, %s), want (%q, %d, %s)",
					tt.currency, tt.limit, describe(tt.categories),
					cur, limit, describe(cats),
					tt.wantCurrency, tt.wantLimit, describe(tt.wantCategories))
			}
		})
	}
}

func TestAddSpend(t *testing.T) {
	tests := []struct {
		name     string
		totals   map[string]int
		category string
		cents    int
		want     map[string]int
	}{
		{name: "nil map", totals: nil, category: "Food", cents: 1250,
			want: map[string]int{"Food": 1250}},
		{name: "new category", totals: map[string]int{"Rent": 120000}, category: "Food", cents: 500,
			want: map[string]int{"Rent": 120000, "Food": 500}},
		{name: "existing category", totals: map[string]int{"Food": 1000}, category: "Food", cents: 250,
			want: map[string]int{"Food": 1250}},
		{name: "refund lowers the total", totals: map[string]int{"Food": 1000}, category: "Food", cents: -300,
			want: map[string]int{"Food": 700}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := fmt.Sprint(tt.totals)
			got := AddSpend(tt.totals, tt.category, tt.cents)
			if fmt.Sprint(got) != fmt.Sprint(tt.want) || got == nil {
				t.Fatalf("AddSpend(%s, %q, %d) = %v, want %v", before, tt.category, tt.cents, got, tt.want)
			}
			if tt.totals != nil && fmt.Sprint(tt.totals) != fmt.Sprint(tt.want) {
				t.Errorf("AddSpend(%s, %q, %d) did not update the map it was given: it is now %v, want %v",
					before, tt.category, tt.cents, tt.totals, tt.want)
			}
		})
	}
}
