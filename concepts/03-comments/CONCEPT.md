# Comments

## In plain English

Comments are notes for humans; the computer ignores them. Go has a strong habit: every
**exported** name (anything starting with a capital letter, which other packages can use)
gets a comment right above it that starts with that name. Tools like `go doc`, your editor's
hover pop-ups and pkg.go.dev turn these comments into the package's documentation.

```go
// This is a line comment.

/* This is a block comment.
   Go code mostly uses // even for several lines. */
```

## Doc comments

A **doc comment** is a comment placed *directly* above a declaration, with no blank line in
between. It should be a full sentence that starts with the name it describes:

```go
// Package budget tracks monthly spending limits.
package budget

// DefaultLimitCents is the monthly limit used when none is set.
const DefaultLimitCents = 50000

// Budget is a spending limit for one category.
type Budget struct {
    Category   string
    LimitCents int
}

// Remaining returns how many cents are left to spend in b.
func Remaining(b Budget, spentCents int) int { ... }
```

These are **not** doc comments:

```go
// remaining returns ...      <- doesn't start with the name "Remaining"
func Remaining(...)

// Remaining returns ...
                               <- the blank line detaches the comment
func Remaining(...)
```

Unexported names (lowercase) don't need doc comments, but a short one still helps.
Comments inside a function should explain **why**, not repeat **what** the code says:

```go
total += fee   // add the fee           <- useless, the code already says that
total += fee   // the bank charges the fee even on refunds   <- useful
```

> **Coming from Ruby?** Go uses `//` instead of `#`, and doc comments are plain sentences:
> there are no YARD-style `@param` / `@return` tags.

## Your task

`exercise.go` declares four exported names: `CurrencySymbol`, `Transaction`, `FormatAmount`
and `Describe`. Money is in integer **cents**.

**Problem 1.** Give every exported name in `exercise.go` a proper doc comment: directly
above it, starting with the name followed by a space, e.g. `// FormatAmount returns ...`.
Some have no comment, and some have comments that don't start with the name. The tests
read your source file to check this. (Keep the package comment at the top as it is.)

**Problem 2.** Implement the two functions:

- `FormatAmount(cents int) string` formats cents as dollars using `CurrencySymbol`:
  `1234` → `"$12.34"`, `5` → `"$0.05"`, `0` → `"$0.00"`, `-350` → `"-$3.50"`
  (the minus sign goes before the symbol).
- `Describe(t Transaction) string` returns the description, a colon and a space, then the
  formatted amount: `Transaction{Description: "Coffee", AmountCents: 350}` → `"Coffee: $3.50"`.

Hint: `fmt.Sprintf("%d.%02d", 3, 5)` gives `"3.05"` (`%02d` pads to two digits with zeros).

Run: `go test -v ./concepts/03-comments/`
