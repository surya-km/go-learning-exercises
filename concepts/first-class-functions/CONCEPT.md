# First-Class Functions

## In plain English

In Go, a function is a **value**, just like a number or a string. You can store it in a
variable, pass it into another function, or return it from a function. That lets you
write one general piece of code ("keep the transactions that pass *some* test") and plug
in the specific behaviour ("...the ones over $100") from outside.

## Functions as values

```go
double := func(n int) int { return n * 2 } // an anonymous function stored in a variable
double(21)                                 // 42
```

A function's **type** is written from its parameters and results: `double` has type
`func(int) int`.

## Functions as parameters

```go
func apply(amounts []int, f func(int) int) []int {
    out := []int{}
    for _, a := range amounts {
        out = append(out, f(a))
    }
    return out
}

apply([]int{100, 200}, double) // [200 400]
```

## Functions as return values, and closures

A function created inside another function can use the outer function's variables.
It "closes over" them, so it's called a **closure**. The variable stays alive as long
as the closure does:

```go
func addFee(feeCents int) func(int) int {
    return func(amount int) int {
        return amount + feeCents // feeCents is remembered
    }
}

withFee := addFee(150)
withFee(1000) // 1150
```

A closure can also **change** the variable it captured, so it can keep state between calls:

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

next := counter()
next() // 1
next() // 2
```

> **Coming from Ruby?** Go function values are like lambdas/procs, and closures capture
> variables the same way blocks do. There are no blocks or `yield`: you pass the
> function as a normal argument.

## Your task

`Transaction` is already defined for you in `types.go`.

**Problem 1: `Filter(txns []Transaction, keep func(Transaction) bool) []Transaction`**

Return a new slice with only the transactions for which `keep` returns `true`, in their
original order. With no matches (or no transactions), return an empty, non-nil slice.
Don't modify the input.

```go
bigOnes := Filter(txns, func(t Transaction) bool { return t.AmountCents > 10000 })
```

**Problem 2: `SpendingTracker(limitCents int) func(amountCents int) (remaining int, ok bool)`**

Return a closure that tracks spending against a budget. Each time it's called with an
amount:
- If the amount fits in what's left of the budget, subtract it and return the new
  remaining amount and `true`.
- If it would go over the budget, **don't** subtract it. Return the unchanged remaining
  amount and `false`.

Each tracker keeps its own separate budget.

```go
spend := SpendingTracker(5000)
spend(2000) // 3000, true
spend(4000) // 3000, false  (would go over)
spend(3000) // 0, true      (spending exactly the rest is fine)
```

Run: `go test -v ./concepts/first-class-functions/`
