package errorwrapping

import (
	"errors"
	"fmt"
)

// Payment is a request to charge a card.
type Payment struct {
	ID          string
	CardID      string
	AmountCents int64
}

// Sentinel errors.
var (
	ErrInvalidAmount = errors.New("invalid amount")
	ErrCardExpired   = errors.New("card expired")
	ErrUnknownCard   = errors.New("unknown card")
)

// DeclinedError means the bank refused the charge. Code says why,
// for example "insufficient_funds" or "suspected_fraud".
type DeclinedError struct {
	Code        string
	AmountCents int64
}

func (e *DeclinedError) Error() string {
	return fmt.Sprintf("card declined (%s) for %d cents", e.Code, e.AmountCents)
}

// ChargeCard is a fake payment gateway. Its behaviour depends on the card ID:
//
//	"card-ok"       succeeds
//	"card-expired"  fails with an error wrapping ErrCardExpired
//	"card-poor"     fails with an error wrapping *DeclinedError{Code: "insufficient_funds"}
//	"card-fraud"    fails with an error wrapping *DeclinedError{Code: "suspected_fraud"}
//	anything else   fails with an error wrapping ErrUnknownCard
func ChargeCard(cardID string, cents int64) error {
	switch cardID {
	case "card-ok":
		return nil
	case "card-expired":
		return fmt.Errorf("gateway: card %s: %w", cardID, ErrCardExpired)
	case "card-poor":
		return fmt.Errorf("gateway: card %s: %w", cardID, &DeclinedError{Code: "insufficient_funds", AmountCents: cents})
	case "card-fraud":
		return fmt.Errorf("gateway: card %s: %w", cardID, &DeclinedError{Code: "suspected_fraud", AmountCents: cents})
	default:
		return fmt.Errorf("gateway: %w %q", ErrUnknownCard, cardID)
	}
}
