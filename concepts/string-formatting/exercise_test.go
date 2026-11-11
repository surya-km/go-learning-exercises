package stringformatting

import "testing"

func TestStatementLine(t *testing.T) {
	tests := []struct {
		description string
		cents       int
		want        string
	}{
		{description: "Coffee", cents: 450, want: "Coffee                    $4.50"},
		{description: "Rent", cents: 123400, want: "Rent                   $1234.00"},
		{description: "Refund", cents: -1205, want: "Refund                  -$12.05"},
		{description: "Gum", cents: 7, want: "Gum                       $0.07"},
		{description: "Exactly twenty chars", cents: 100, want: "Exactly twenty chars      $1.00"},
		{description: "Payroll", cents: 1234567890, want: "Payroll              $12345678.90"},
		{description: "", cents: 0, want: "                          $0.00"},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := StatementLine(tt.description, tt.cents)
			if got != tt.want {
				t.Errorf("StatementLine(%q, %d)\n got %q\nwant %q", tt.description, tt.cents, got, tt.want)
			}
		})
	}
}

func TestSpendingShare(t *testing.T) {
	tests := []struct {
		category     string
		spent, total int
		want         string
	}{
		{category: "Groceries", spent: 2500, total: 10000, want: "Groceries: 25.0%"},
		{category: "Rent", spent: 2, total: 3, want: "Rent: 66.7%"},
		{category: "Travel", spent: 1, total: 3, want: "Travel: 33.3%"},
		{category: "Everything", spent: 5000, total: 5000, want: "Everything: 100.0%"},
		{category: "Fun", spent: 0, total: 5000, want: "Fun: 0.0%"},
		{category: "Misc", spent: 100, total: 0, want: "Misc: 0.0%"},
	}

	for _, tt := range tests {
		t.Run(tt.category, func(t *testing.T) {
			got := SpendingShare(tt.category, tt.spent, tt.total)
			if got != tt.want {
				t.Errorf("SpendingShare(%q, %d, %d) = %q, want %q", tt.category, tt.spent, tt.total, got, tt.want)
			}
		})
	}
}
