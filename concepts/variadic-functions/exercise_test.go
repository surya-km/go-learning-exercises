package variadicfunctions

import "testing"

func TestMaxAmount(t *testing.T) {
	tests := []struct {
		name    string
		amounts []int
		want    int
		wantOK  bool
	}{
		{name: "largest in the middle", amounts: []int{300, 1200, 450}, want: 1200, wantOK: true},
		{name: "largest first", amounts: []int{900, 5, 10}, want: 900, wantOK: true},
		{name: "largest last", amounts: []int{1, 2, 3}, want: 3, wantOK: true},
		{name: "single amount", amounts: []int{42}, want: 42, wantOK: true},
		{name: "all negative", amounts: []int{-50, -10, -30}, want: -10, wantOK: true},
		{name: "empty slice", amounts: []int{}, want: 0, wantOK: false},
		{name: "nil slice", amounts: nil, want: 0, wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := MaxAmount(tt.amounts...) // pass a slice with "..."
			if got != tt.want || ok != tt.wantOK {
				t.Errorf("MaxAmount(%v...) = (%d, %t), want (%d, %t)", tt.amounts, got, ok, tt.want, tt.wantOK)
			}
		})
	}
}

func TestMaxAmountSeparateArguments(t *testing.T) {
	if got, ok := MaxAmount(300, 1200, 450); got != 1200 || !ok {
		t.Errorf("MaxAmount(300, 1200, 450) = (%d, %t), want (1200, true)", got, ok)
	}
	if got, ok := MaxAmount(); got != 0 || ok {
		t.Errorf("MaxAmount() = (%d, %t), want (0, false)", got, ok)
	}
}

func TestAddToBill(t *testing.T) {
	tests := []struct {
		name      string
		billCents int
		items     []int
		want      int
	}{
		{name: "adds two items", billCents: 1000, items: []int{250, 399}, want: 1649},
		{name: "starts from zero", billCents: 0, items: []int{100, 200}, want: 300},
		{name: "includes a refund", billCents: 1000, items: []int{500, -200}, want: 1300},
		{name: "no items", billCents: 1000, items: nil, want: 1000},
		{name: "one item", billCents: 50, items: []int{25}, want: 75},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddToBill(tt.billCents, tt.items...); got != tt.want {
				t.Errorf("AddToBill(%d, %v...) = %d, want %d", tt.billCents, tt.items, got, tt.want)
			}
		})
	}
}

func TestAddToBillSeparateArguments(t *testing.T) {
	if got := AddToBill(1000, 250, 399); got != 1649 {
		t.Errorf("AddToBill(1000, 250, 399) = %d, want 1649", got)
	}
	if got := AddToBill(1000); got != 1000 {
		t.Errorf("AddToBill(1000) = %d, want 1000", got)
	}
}
