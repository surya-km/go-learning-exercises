# Conditionals: if

## In plain English

`if` lets your code make a decision: "if this is true, do that; otherwise, do something
else". You can chain several checks with `else if`. Go checks them **from top to bottom**
and runs only the first one that matches.

```go
if balance < 0 {
    fmt.Println("overdrawn")
} else if balance == 0 {
    fmt.Println("empty")
} else {
    fmt.Println("in credit")
}
```

The condition must be a `bool`. There are no parentheses around it, but the curly braces
are **always** required, even for one line.

## Return early

Idiomatic Go avoids deep nesting and `else` after a `return`. Handle the special cases
first, return, and let the "normal" case run at the bottom:

```go
func shippingCents(orderCents int) int {
    if orderCents >= 5000 {
        return 0         // free shipping over $50
    }
    return 499
}
```

## if with a short statement

You can run a short statement before the condition. The variable only exists inside the
`if` / `else` blocks:

```go
if remaining := budget - spent; remaining < 0 {
    fmt.Println("over by", -remaining)
}
// remaining doesn't exist here
```

> **Coming from Ruby?** No `unless`, no trailing `return x if y`, and no truthiness:
> `if name` won't compile; write `if name != ""`. It's `else if`, not `elsif`.

## Your task

Amounts are in integer **cents**.

**Problem 1.** Implement `OverdraftFee(balanceCents int) int`. After a payment the bank
looks at the new balance:

| Balance | Fee |
|---|---|
| `0` or more | `0` |
| overdrawn by `$5.00` or less (`-1` to `-500`) | `0` (a small grace amount) |
| overdrawn by more than `$5.00` (`-501` or lower) | `2500` |

**Problem 2.** Implement `BudgetStatus(spentCents, budgetCents int) string`. Return:

1. `"no budget"` if `budgetCents` is `0` or negative (check this first),
2. `"over"` if `spentCents` is more than `budgetCents`,
3. `"warning"` if `spentCents` is at least **80%** of `budgetCents`,
4. `"ok"` otherwise.

Tip: avoid fractions. "spent is at least 80% of budget" is the same as
`spentCents*100 >= budgetCents*80`.

Run: `go test -v ./concepts/conditionals-if/`
