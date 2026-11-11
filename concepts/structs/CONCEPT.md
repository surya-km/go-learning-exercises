# Structs

## In plain English

A **struct** bundles a few related values into one thing with named fields, like a row
on a form: an account has an owner, a currency and a balance. Instead of passing three
loose variables around, you pass one `Account`.

```go
type Account struct {
    Owner        string
    Currency     string
    BalanceCents int64
}
```

> **Coming from Ruby?** A struct is like a `Struct.new(:owner, :currency, :balance)`,
> or a plain class with only attributes. There are no constructors built in: you write an
> ordinary function (usually called `NewXxx`) when you want one.

## Creating struct values

```go
a := Account{Owner: "Ana", Currency: "USD", BalanceCents: 1500} // named fields (preferred)
b := Account{Owner: "Ben"}   // missing fields get their zero value: "" and 0
var c Account                // every field is the zero value
```

## Reading and changing fields

```go
fmt.Println(a.Owner)    // "Ana"
a.BalanceCents += 250   // fields are just variables
```

## Comparing structs

If every field is comparable (strings, numbers, bools...), you can compare whole structs with `==`.
Two structs are equal when **every** field is equal:

```go
Account{Owner: "Ana"} == Account{Owner: "Ana"} // true
```

## Returning a struct from a function

```go
func NewTransaction(desc string, cents int64) Transaction {
    return Transaction{Description: desc, AmountCents: cents}
}
```

## Your task

The `Account`, `Transaction` and `Summary` structs are already declared in `types.go`.
Don't change them. Write your code in `exercise.go`.

**Problem 1: `NewAccount(owner string, openingCents int64) Account`**

Build a new account:
- `Owner` is `owner` with surrounding spaces removed (`strings.TrimSpace`).
- `Currency` is always `"USD"`.
- `BalanceCents` is `openingCents`, but a negative opening amount becomes `0`.

```go
NewAccount("  Ana ", 1500) // Account{Owner: "Ana", Currency: "USD", BalanceCents: 1500}
NewAccount("Ben", -20)     // Account{Owner: "Ben", Currency: "USD", BalanceCents: 0}
```

**Problem 2: `Summarize(txns []Transaction) Summary`**

Walk through the transactions and fill in a `Summary`:
- `Count` is the number of transactions.
- `IncomeCents` is the sum of all **positive** amounts.
- `SpendingCents` is the sum of all **negative** amounts, reported as a **positive** number.
- A zero amount counts towards `Count` only.
- An empty or nil slice gives the zero `Summary{}`.

```go
Summarize([]Transaction{
    {Description: "Salary", AmountCents: 300000},
    {Description: "Rent", AmountCents: -120000},
    {Description: "Coffee", AmountCents: -450},
})
// Summary{Count: 3, IncomeCents: 300000, SpendingCents: 120450}
```

Run: `go test -v ./concepts/structs/`
