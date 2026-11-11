# Functions

## In plain English

A **function** is a named, reusable recipe. You give it some inputs (parameters), it does
some work, and it hands back a result. In Go you must write down the **type** of every
input and of the result, so the compiler can check that you call it correctly.

```go
func add(a int, b int) int {
    return a + b
}

total := add(2, 3) // 5
```

Parameters that share a type can share the type name:

```go
func add(a, b int) int { return a + b }
```

A function with no result just leaves the return type off:

```go
func logSpend(cents int) {
    fmt.Println("spent", cents, "cents")
}
```

> **Coming from Ruby?** Go has no implicit return: you must write `return` explicitly.
> Every parameter and result needs a type, and you can't skip arguments or give them
> default values.

## Functions calling functions

Small functions that call each other are easier to read and test than one big function:

```go
func taxCents(amountCents int) int {
    return amountCents * 8 / 100
}

func priceWithTax(amountCents int) int {
    return amountCents + taxCents(amountCents) // reuse, don't repeat
}
```

## Names: capital letter = public

A function whose name starts with a **capital letter** (`Greeting`) can be used from
other packages. A lowercase name (`taxCents`) is private to its own package. You'll see
this in detail in the packages exercise.

## Integer division

Dividing two `int`s throws away the remainder: `7 / 2 == 3`, and `1999 * 15 / 100 == 299`
(299.85 rounded *down*). That's fine for this exercise.

## Your task

Implement in `exercise.go`:

**Problem 1: `Greeting(name string) string`**

Return `"Hello, <name>!"`. If `name` is empty, use `"friend"` instead:
`Greeting("Ada") == "Hello, Ada!"`, `Greeting("") == "Hello, friend!"`.

**Problem 2: tipping**

- **`TipCents(billCents, percent int) int`**: the tip on a bill, `billCents * percent / 100`
  (rounded down by integer division). A `percent` of 0 or less means no tip, so return `0`.
- **`TotalWithTip(billCents, percent int) int`**: the bill plus the tip. **Call `TipCents`**
  inside it instead of repeating the maths.

| Call | Result |
|---|---|
| `TipCents(2000, 15)` | `300` |
| `TipCents(1999, 15)` | `299` (299.85 rounded down) |
| `TipCents(2000, -5)` | `0` |
| `TotalWithTip(2000, 20)` | `2400` |
| `TotalWithTip(2000, 0)` | `2000` |

Run: `go test -v ./concepts/functions/`
