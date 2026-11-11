package stringers

import (
	"fmt"
	"testing"
)

// These lines fail to compile if a type has no String() string method.
var (
	_ fmt.Stringer = Money(0)
	_ fmt.Stringer = Transaction{}
	_ fmt.Stringer = Statement{}
)

func TestMoneyString(t *testing.T) {
	tests := []struct {
		money Money
		want  string
	}{
		{money: 1999, want: "$19.99"},
		{money: 5, want: "$0.05"},
		{money: 0, want: "$0.00"},
		{money: 100000, want: "$1000.00"},
		{money: -450, want: "-$4.50"},
		{money: -1, want: "-$0.01"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := fmt.Sprint(tt.money); got != tt.want {
				t.Errorf("fmt.Sprint(Money(%d)) = %q, want %q", int64(tt.money), got, tt.want)
			}
			if got := fmt.Sprintf("total: %v", tt.money); got != "total: "+tt.want {
				t.Errorf("fmt.Sprintf(\"total: %%v\", Money(%d)) = %q, want %q", int64(tt.money), got, "total: "+tt.want)
			}
		})
	}
}

func TestTransactionString(t *testing.T) {
	tests := []struct {
		txn  Transaction
		want string
	}{
		{txn: Transaction{Date: "2024-03-01", Description: "Coffee", Amount: -450}, want: "2024-03-01 Coffee -$4.50"},
		{txn: Transaction{Date: "2024-03-15", Description: "Salary", Amount: 300000}, want: "2024-03-15 Salary $3000.00"},
		{txn: Transaction{Date: "2024-12-31", Description: "Interest", Amount: 3}, want: "2024-12-31 Interest $0.03"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := fmt.Sprintf("%s", tt.txn); got != tt.want {
				t.Errorf("fmt.Sprintf(\"%%s\", %#v) = %q, want %q", tt.txn, got, tt.want)
			}
		})
	}
}

func TestStatementString(t *testing.T) {
	tests := []struct {
		name string
		stmt Statement
		want string
	}{
		{
			name: "two transactions",
			stmt: Statement{Owner: "Ana", Transactions: []Transaction{
				{Date: "2024-03-01", Description: "Salary", Amount: 300000},
				{Date: "2024-03-02", Description: "Coffee", Amount: -450},
			}},
			want: "Statement for Ana (2 transactions)\n" +
				"2024-03-01 Salary $3000.00\n" +
				"2024-03-02 Coffee -$4.50\n",
		},
		{
			name: "one transaction is singular",
			stmt: Statement{Owner: "Ben", Transactions: []Transaction{
				{Date: "2024-04-10", Description: "Rent", Amount: -120000},
			}},
			want: "Statement for Ben (1 transaction)\n" +
				"2024-04-10 Rent -$1200.00\n",
		},
		{
			name: "no transactions",
			stmt: Statement{Owner: "Cara"},
			want: "Statement for Cara (0 transactions)\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fmt.Sprint(tt.stmt); got != tt.want {
				t.Errorf("fmt.Sprint(%#v) =\n%q\nwant\n%q", tt.stmt, got, tt.want)
			}
		})
	}
}
