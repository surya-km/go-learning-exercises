// Package basics practices the basics concept. Read CONCEPT.md first.
package basics

import "fmt"

// Greeting returns "Hello, <name>!", or "Hello, friend!" when name is empty.
func Greeting(name string) string {
	if name == "" {
		name = "friend"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

// SplitBill splits totalCents evenly between people.
// It returns the amount each person pays and the leftover cents.
// If people <= 0, it returns 0 and totalCents.
func SplitBill(totalCents, people int) (each, leftover int) {
	if people <= 0 {
		return 0, totalCents
	}
	return totalCents / people, totalCents % people
}
