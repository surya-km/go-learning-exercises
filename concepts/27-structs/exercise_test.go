package structs

import "testing"

func TestNewAccount(t *testing.T) {
	tests := []struct {
		name    string
		owner   string
		opening int64
		want    Account
	}{
		{name: "simple", owner: "Ana", opening: 1500, want: Account{Owner: "Ana", Currency: "USD", BalanceCents: 1500}},
		{name: "trims spaces", owner: "  Ana Lopez ", opening: 99, want: Account{Owner: "Ana Lopez", Currency: "USD", BalanceCents: 99}},
		{name: "zero opening", owner: "Ben", opening: 0, want: Account{Owner: "Ben", Currency: "USD", BalanceCents: 0}},
		{name: "negative opening becomes zero", owner: "Cara", opening: -20, want: Account{Owner: "Cara", Currency: "USD", BalanceCents: 0}},
		{name: "empty owner still gets USD", owner: "", opening: 5, want: Account{Owner: "", Currency: "USD", BalanceCents: 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAccount(tt.owner, tt.opening)
			if got != tt.want {
				t.Errorf("NewAccount(%q, %d) = %+v, want %+v", tt.owner, tt.opening, got, tt.want)
			}
		})
	}
}

func TestSummarize(t *testing.T) {
	tests := []struct {
		name string
		txns []Transaction
		want Summary
	}{
		{
			name: "mixed income and spending",
			txns: []Transaction{
				{Description: "Salary", AmountCents: 300000},
				{Description: "Rent", AmountCents: -120000},
				{Description: "Coffee", AmountCents: -450},
			},
			want: Summary{Count: 3, IncomeCents: 300000, SpendingCents: 120450},
		},
		{
			name: "only spending",
			txns: []Transaction{{Description: "Lunch", AmountCents: -1200}, {Description: "Bus", AmountCents: -275}},
			want: Summary{Count: 2, SpendingCents: 1475},
		},
		{
			name: "only income",
			txns: []Transaction{{Description: "Refund", AmountCents: 999}},
			want: Summary{Count: 1, IncomeCents: 999},
		},
		{
			name: "zero amount only counts",
			txns: []Transaction{{Description: "Card check", AmountCents: 0}, {Description: "Gift", AmountCents: 500}},
			want: Summary{Count: 2, IncomeCents: 500},
		},
		{name: "empty slice", txns: []Transaction{}, want: Summary{}},
		{name: "nil slice", txns: nil, want: Summary{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Summarize(tt.txns)
			if got != tt.want {
				t.Errorf("Summarize(%+v) = %+v, want %+v", tt.txns, got, tt.want)
			}
		})
	}
}
