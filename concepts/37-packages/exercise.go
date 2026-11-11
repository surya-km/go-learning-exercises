// Package packages practices the packages concept. Read CONCEPT.md first.
package packages

import "go-learning-exercises/concepts/37-packages/money"

// ReceiptLine returns "<name>: <price>", e.g. "Coffee: $4.50".
// It should use money.Format.
func ReceiptLine(item Item) string {
	// TODO: implement
	return ""
}

// Total returns the sum of all item prices.
func Total(items []Item) money.Cents {
	// TODO: implement
	return 0
}
