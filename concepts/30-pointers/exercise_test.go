package pointers

import (
	"errors"
	"testing"
)

func TestDeposit(t *testing.T) {
	tests := []struct {
		name    string
		start   int64
		deposit int64
		want    int64
	}{
		{name: "into empty account", start: 0, deposit: 500, want: 500},
		{name: "adds to balance", start: 1000, deposit: 250, want: 1250},
		{name: "zero is ignored", start: 1000, deposit: 0, want: 1000},
		{name: "negative is ignored", start: 1000, deposit: -300, want: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acct := &Account{Owner: "Ana", BalanceCents: tt.start}
			acct.Deposit(tt.deposit)
			if acct.BalanceCents != tt.want {
				t.Errorf("balance %d, Deposit(%d): balance = %d, want %d", tt.start, tt.deposit, acct.BalanceCents, tt.want)
			}
		})
	}
}

func TestDepositOnValueVariable(t *testing.T) {
	// Not a pointer variable: Go takes its address for you because Deposit
	// has a pointer receiver, so the change must stick.
	acct := Account{Owner: "Ben", BalanceCents: 100}
	acct.Deposit(50)
	if acct.BalanceCents != 150 {
		t.Errorf("balance 100, Deposit(50) on a value variable: balance = %d, want 150", acct.BalanceCents)
	}
}

func TestWithdraw(t *testing.T) {
	tests := []struct {
		name     string
		start    int64
		withdraw int64
		want     int64
		wantErr  error
	}{
		{name: "part of balance", start: 1000, withdraw: 400, want: 600},
		{name: "whole balance", start: 1000, withdraw: 1000, want: 0},
		{name: "too much", start: 1000, withdraw: 1001, want: 1000, wantErr: ErrInsufficientFunds},
		{name: "from empty account", start: 0, withdraw: 1, want: 0, wantErr: ErrInsufficientFunds},
		{name: "zero is ignored", start: 1000, withdraw: 0, want: 1000},
		{name: "negative is ignored", start: 1000, withdraw: -5, want: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acct := &Account{Owner: "Ana", BalanceCents: tt.start}
			err := acct.Withdraw(tt.withdraw)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("balance %d, Withdraw(%d) error = %v, want %v", tt.start, tt.withdraw, err, tt.wantErr)
			}
			if acct.BalanceCents != tt.want {
				t.Errorf("balance %d, Withdraw(%d): balance = %d, want %d", tt.start, tt.withdraw, acct.BalanceCents, tt.want)
			}
		})
	}
}

func TestTransfer(t *testing.T) {
	tests := []struct {
		name      string
		fromStart int64
		toStart   int64
		cents     int64
		wantFrom  int64
		wantTo    int64
		wantErr   error
	}{
		{name: "simple transfer", fromStart: 5000, toStart: 100, cents: 1500, wantFrom: 3500, wantTo: 1600},
		{name: "everything", fromStart: 5000, toStart: 0, cents: 5000, wantFrom: 0, wantTo: 5000},
		{name: "not enough money", fromStart: 500, toStart: 100, cents: 501, wantFrom: 500, wantTo: 100, wantErr: ErrInsufficientFunds},
		{name: "zero does nothing", fromStart: 500, toStart: 100, cents: 0, wantFrom: 500, wantTo: 100},
		{name: "negative does nothing", fromStart: 500, toStart: 100, cents: -50, wantFrom: 500, wantTo: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from := &Account{Owner: "Ana", BalanceCents: tt.fromStart}
			to := &Account{Owner: "Ben", BalanceCents: tt.toStart}
			err := Transfer(from, to, tt.cents)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Transfer(from=%d, to=%d, %d) error = %v, want %v", tt.fromStart, tt.toStart, tt.cents, err, tt.wantErr)
			}
			if from.BalanceCents != tt.wantFrom || to.BalanceCents != tt.wantTo {
				t.Errorf("Transfer(from=%d, to=%d, %d): balances = (%d, %d), want (%d, %d)",
					tt.fromStart, tt.toStart, tt.cents, from.BalanceCents, to.BalanceCents, tt.wantFrom, tt.wantTo)
			}
		})
	}
}

func TestTransferNilAccounts(t *testing.T) {
	t.Run("nil from", func(t *testing.T) {
		to := &Account{Owner: "Ben", BalanceCents: 100}
		err := Transfer(nil, to, 50)
		if !errors.Is(err, ErrNilAccount) {
			t.Errorf("Transfer(nil, to, 50) error = %v, want %v", err, ErrNilAccount)
		}
		if to.BalanceCents != 100 {
			t.Errorf("Transfer(nil, to, 50): to balance = %d, want 100", to.BalanceCents)
		}
	})

	t.Run("nil to", func(t *testing.T) {
		from := &Account{Owner: "Ana", BalanceCents: 100}
		err := Transfer(from, nil, 50)
		if !errors.Is(err, ErrNilAccount) {
			t.Errorf("Transfer(from, nil, 50) error = %v, want %v", err, ErrNilAccount)
		}
		if from.BalanceCents != 100 {
			t.Errorf("Transfer(from, nil, 50): from balance = %d, want 100", from.BalanceCents)
		}
	})

	t.Run("both nil", func(t *testing.T) {
		if err := Transfer(nil, nil, 50); !errors.Is(err, ErrNilAccount) {
			t.Errorf("Transfer(nil, nil, 50) error = %v, want %v", err, ErrNilAccount)
		}
	})
}
