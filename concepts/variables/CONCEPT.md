# Variables

## In plain English

A variable is a labelled box that holds a value you can change later. In Go, every box
has a fixed **type** (a box for whole numbers can never hold text), and the compiler
checks this before the program runs. Go also refuses to compile if you create a variable
and never use it.

```go
var balance int = 1000   // long form: name, type, value
var owner = "Ada"        // type is inferred from the value (string)
fee := 250               // short form ":=" declares AND infers (only inside functions)

balance = balance - fee  // "=" changes an existing variable
balance -= fee           // the same thing, shorter
// balance = "lots"      // compile error: balance is an int forever
```

`:=` makes a **new** variable; `=` updates one that already exists. Using `:=` twice for
the same name in the same block is an error.

## Several at once

```go
var (
    income int = 300000
    rent   int = 120000
)

a, b := 1, 2
a, b = b, a      // swap in one line: a is now 2, b is 1
```

## Named results

A function can name its return values. They start out as variables you can assign to:

```go
func minMax(x, y int) (lo, hi int) {
    lo, hi = x, y
    if lo > hi {
        lo, hi = hi, lo
    }
    return lo, hi
}
```

> **Coming from Ruby?** In Ruby, `balance = "lots"` silently changes the type. In Go,
> types never change. And `x := 5` vs `x = 5` matters: the first creates `x`, the second
> needs `x` to exist already.

## Your task

Money is in integer **cents**. Use well-named variables for intermediate steps rather
than one giant expression.

**Problem 1.** Implement
`MonthEnd(startCents, incomeCents, rentCents, groceriesCents int) (endCents, spentCents int)`.
It returns the balance at the end of the month (start + income − rent − groceries) and the
total spent (rent + groceries). The end balance may be negative.

- `MonthEnd(50000, 300000, 120000, 40000)` returns `190000, 160000`.

**Problem 2.** Implement `Transfer(fromCents, toCents, amountCents int) (newFrom, newTo int)`.
It moves money from one account to another and returns both new balances.

- Move `amountCents`, but **never more than `fromCents` holds**: if there isn't enough,
  move everything that's there.
- If `amountCents` is `0` or negative, nothing moves.
- `Transfer(1000, 500, 300)` returns `700, 800`.
- `Transfer(1000, 500, 5000)` returns `0, 1500`.

Run: `go test -v ./concepts/variables/`
