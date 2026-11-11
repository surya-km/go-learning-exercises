# Constants

## In plain English

A constant is a name for a value that **never changes** while the program runs, like the
number of months in a year. Go works out constants before the program even starts, so
you can't set one from a calculation done at runtime. Giving magic numbers a name makes
code easier to read and stops typos.

```go
const MonthsInYear = 12            // untyped constant: fits wherever a number fits
const FreeShippingCents int = 5000 // typed constant: it's an int, full stop

// MonthsInYear = 13                // compile error: cannot assign to a constant
```

Group related constants with parentheses:

```go
const (
    WeeksPerYear  = 52
    MonthsPerYear = 12
)
```

## Your own type + `iota` = an "enum"

Go has no `enum` keyword. Instead you make a small named type and list its values with
`iota`. Inside a `const ( ... )` block, `iota` starts at `0` and goes up by one on each line.
Lines with nothing after the name repeat the previous line's type and expression:

```go
type Size int

const (
    Small  Size = iota // 0
    Medium             // 1
    Large              // 2
)

var s Size = Medium
```

Because `Size` is its own type, a function that takes a `Size` documents exactly what it
expects. A common trick is to make the **first** value (`0`) mean "unknown", so a variable
you forgot to set doesn't accidentally look valid.

> **Coming from Ruby?** Ruby constants are just capitalised names that *warn* when you
> reassign them. Go constants truly can't change. Where Ruby uses symbols like `:monthly`,
> Go usually uses a typed constant like `Monthly`.

## Your task

Money is in integer **cents**.

**Problem 1.** In `exercise.go` there is a `Frequency` type and a `const` block where every
value is a placeholder `0`. Rewrite that block with `iota` so the values are, in order:

| Constant | Value |
|---|---|
| `Unknown` | `0` |
| `Weekly` | `1` |
| `Fortnightly` | `2` |
| `Monthly` | `3` |
| `Yearly` | `4` |

**Problem 2.** Implement `YearlyCostCents(amountCents int, f Frequency) int`. It returns how
much a repeating payment costs over a year. Declare constants for the multipliers (don't
scatter bare `52`s in your code):

- `Weekly`: 52 payments a year
- `Fortnightly`: 26 payments a year
- `Monthly`: 12 payments a year
- `Yearly`: 1 payment a year
- `Unknown` or any other value: return `0`

For example, `YearlyCostCents(1099, Monthly)` returns `13188`.

Run: `go test -v ./concepts/04-constants/`
