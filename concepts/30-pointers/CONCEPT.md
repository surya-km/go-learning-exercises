# Pointers

## In plain English

A **pointer** is the address of a value, not the value itself. Think of handing someone
your house address instead of a photo of your house: with the address they can go and
actually repaint the front door. In Go, passing a normal value hands over a *copy*;
passing a pointer lets the function change the *original*.

```go
n := 10
p := &n      // & means "address of": p is a *int (pointer to int)
*p = 20      // * means "the value at this address"
fmt.Println(n) // 20: we changed n through the pointer
```

## Copies vs pointers

```go
func addTen(x int)   { x += 10 }   // changes a copy; caller sees nothing
func addTenP(x *int) { *x += 10 }  // changes the caller's variable

n := 1
addTen(n)    // n is still 1
addTenP(&n)  // n is now 11
```

## Pointer receivers: methods that change the value

A method with a value receiver `(a Account)` works on a copy. To **change** the
account, use a pointer receiver `(a *Account)`:

```go
func (a *Account) Deposit(cents int64) {
    a.BalanceCents += cents   // Go lets you write a.Field instead of (*a).Field
}

acct := Account{}
acct.Deposit(500)   // Go automatically passes &acct for you
```

## nil: a pointer to nothing

The zero value of a pointer is `nil`. Using a nil pointer's fields crashes the program
("nil pointer dereference"), so check first when a pointer might be nil:

```go
func Owner(a *Account) string {
    if a == nil {
        return "unknown"
    }
    return a.Owner
}
```

> **Coming from Ruby?** In Ruby every object is passed by reference, so a method can
> always mutate it. In Go structs are copied unless you pass a pointer. `nil` is like
> Ruby's `nil`, but calling a field on it is a crash instead of a `NoMethodError`.

## Your task

`Account` and `ErrInsufficientFunds` are declared in `types.go`. Implement these in `exercise.go`.

**Problem 1: pointer-receiver methods on `Account`**

- `func (a *Account) Deposit(cents int64)` adds `cents` to `BalanceCents`. Amounts of
  zero or less are ignored (the balance does not change).
- `func (a *Account) Withdraw(cents int64) error` subtracts `cents` from `BalanceCents`
  and returns `nil`. If `cents` is more than the balance, return `ErrInsufficientFunds`
  and leave the balance **unchanged**. Amounts of zero or less are ignored and return `nil`.

**Problem 2: `Transfer(from, to *Account, cents int64) error`**

Moves money between two accounts:
- If either pointer is `nil`, return `ErrNilAccount` and change nothing.
- If `from` doesn't have enough money, return `ErrInsufficientFunds` and change nothing.
- Otherwise take `cents` from `from`, add them to `to`, and return `nil`.
- A `cents` of zero or less does nothing and returns `nil` (after the nil checks).

Run: `go test -v ./concepts/30-pointers/`
