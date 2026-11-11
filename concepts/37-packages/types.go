package packages

import "go-learning-exercises/concepts/37-packages/money"

// Item is one line on a receipt. Its Price uses the Cents type from the
// money package.
type Item struct {
	Name  string
	Price money.Cents
}
