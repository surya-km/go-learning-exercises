package packages

import (
	"testing"

	"go-learning-exercises/concepts/packages/money"
)

func TestReceiptLine(t *testing.T) {
	tests := []struct {
		name string
		item Item
		want string
	}{
		{name: "coffee", item: Item{Name: "Coffee", Price: 450}, want: "Coffee: $4.50"},
		{name: "rent", item: Item{Name: "Rent", Price: 120000}, want: "Rent: $1200.00"},
		{name: "small change", item: Item{Name: "Gum", Price: 5}, want: "Gum: $0.05"},
		{name: "refund", item: Item{Name: "Refund", Price: -999}, want: "Refund: -$9.99"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReceiptLine(tt.item); got != tt.want {
				t.Errorf("ReceiptLine(%+v) = %q, want %q", tt.item, got, tt.want)
			}
		})
	}
}

func TestTotal(t *testing.T) {
	tests := []struct {
		name  string
		items []Item
		want  money.Cents
	}{
		{
			name:  "several items",
			items: []Item{{Name: "Coffee", Price: 450}, {Name: "Bagel", Price: 325}, {Name: "Juice", Price: 599}},
			want:  1374,
		},
		{
			name:  "with a refund",
			items: []Item{{Name: "Shoes", Price: 5000}, {Name: "Refund", Price: -1500}},
			want:  3500,
		},
		{name: "single item", items: []Item{{Name: "Gum", Price: 5}}, want: 5},
		{name: "no items", items: nil, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Total(tt.items); got != tt.want {
				t.Errorf("Total(%+v) = %d, want %d", tt.items, got, tt.want)
			}
		})
	}
}
