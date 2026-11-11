# Methods

## In plain English

A **method** is a function that belongs to a type. Instead of `Total(cart)` you write
`cart.Total()`. It reads like asking the value a question about itself, and it keeps
the code about a type right next to that type.

```go
type Bill struct {
    Name        string
    AmountCents int64
    DueDay      int
}

// (b Bill) is the "receiver": the value the method is called on.
func (b Bill) IsDueBy(day int) bool {
    return b.DueDay <= day
}

rent := Bill{Name: "Rent", AmountCents: 120000, DueDay: 1}
rent.IsDueBy(5) // true
```

> **Coming from Ruby?** The receiver `b` plays the role of `self`, but you name it
> yourself (by convention a short one- or two-letter name based on the type).

## Value receivers

A method with a receiver like `(b Bill)` gets a **copy** of the value. It is perfect for
methods that only *read* fields and compute an answer. (Methods that *change* the value
need a pointer receiver; that's the next concept, `pointers`.)

## Methods on any named type, not just structs

You can add methods to any type you declare in your package, including a slice type:

```go
type Bills []Bill

func (bs Bills) Count() int {
    return len(bs)
}
```

## Your task

`Bill` and `Bills` are declared in `types.go`. Implement these methods in `exercise.go`.
All of them only read data, so they use value receivers.

**Problem 1: `func (b Bill) IsDueBy(day int) bool`**

Returns `true` if the bill's `DueDay` is on or before `day`.

**Problem 1 (continued): `func (b Bill) Label() string`**

Returns the bill name and its amount formatted as dollars with two decimals:
`Bill{Name: "Rent", AmountCents: 120050}.Label()` returns `"Rent: $1200.50"`.
`Bill{Name: "Water", AmountCents: 7}.Label()` returns `"Water: $0.07"`.
(Tip: `fmt.Sprintf("%s: $%d.%02d", name, cents/100, cents%100)`.) Amounts are never negative.

**Problem 2: `func (bs Bills) TotalDueBy(day int) int64`**

Adds up the `AmountCents` of every bill that `IsDueBy(day)`. Reuse your `IsDueBy`
method. An empty list gives `0`.

```go
bills := Bills{
    {Name: "Rent", AmountCents: 120000, DueDay: 1},
    {Name: "Phone", AmountCents: 4500, DueDay: 15},
    {Name: "Gym", AmountCents: 3000, DueDay: 28},
}
bills.TotalDueBy(15) // 124500
```

Run: `go test -v ./concepts/methods/`
