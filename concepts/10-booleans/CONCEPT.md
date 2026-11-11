# Booleans

## In plain English

A boolean (`bool`) is a yes/no value: it is either `true` or `false`. You combine yes/no
answers with three operators: **and** (`&&`), **or** (`||`) and **not** (`!`). This is how
code answers questions like "is the account open *and* is there enough money?".

```go
var open bool = true
hasMoney := false

open && hasMoney   // false: && needs BOTH to be true
open || hasMoney   // true:  || needs AT LEAST ONE to be true
!open              // false: ! flips the value
```

Comparisons like `a > b` or `a == b` produce a `bool`, so you can combine them:

```go
canBuy := price <= balance && !frozen
```

## Return the expression directly

You rarely need `if x { return true } else { return false }`. Just return the expression:

```go
// Instead of this...
if age >= 18 {
    return true
}
return false

// ...write this:
return age >= 18
```

## Short-circuiting

`&&` stops as soon as it sees `false`, and `||` stops as soon as it sees `true`. The
right-hand side is only evaluated if it's needed.

Use parentheses to make mixed `&&` / `||` clear: `a || (b && c)`.

> **Coming from Ruby?** Only `true` and `false` are booleans. Go has no "truthy" values:
> `0`, `""` and `nil` are not `false`, and `if count { ... }` won't compile. Write `if count > 0`.
> Go also has no `and` / `or` / `not` keywords, only `&&`, `||` and `!`.

## Your task

Amounts are in integer **cents**. Try to write each function as a single `return` of a
boolean expression.

**Problem 1.** Implement `CanWithdraw(balanceCents, amountCents int, frozen bool) bool`.
A withdrawal is allowed only when **all** of these are true:

- the account is **not** frozen,
- the amount is more than `0`,
- the amount is not more than the balance (withdrawing the whole balance is fine).

**Problem 2.** Implement `NeedsReview(amountCents int, international, newPayee bool) bool`.
A payment needs a manual fraud review when **either**:

- the amount is `$1,000.00` (`100000` cents) or more, **or**
- it is international **and** going to a new payee.

Run: `go test -v ./concepts/10-booleans/`
