# Stringers

## In plain English

A **Stringer** is any type that knows how to describe itself as text. You give your type
a `String() string` method, and from then on `fmt.Println`, `fmt.Sprint`, `%v` and `%s`
automatically use it. It's like teaching your type to introduce itself nicely, instead
of letting Go print a raw dump of its fields.

```go
type Money int64 // cents

fmt.Println(Money(1999))   // 1999  (no String method: just the number)

func (m Money) String() string {
    return fmt.Sprintf("$%d.%02d", m/100, m%100)
}

fmt.Println(Money(1999))   // $19.99  (fmt found and used String)
```

## Where does it come from?

`fmt` declares a tiny interface:

```go
type Stringer interface {
    String() string
}
```

Because interfaces are satisfied automatically, any type with a `String() string` method
*is* a `fmt.Stringer`, and `fmt` checks for that when printing.

## A trap to avoid

Don't print the value itself with `%v` or `%s` *inside* its own `String` method:

```go
func (m Money) String() string {
    return fmt.Sprintf("%v", m)   // calls String again... forever (stack overflow)
}
```

Format the underlying fields or convert first (`int64(m)`) instead.

> **Coming from Ruby?** `String()` is Go's `to_s`. `fmt.Println` calling it
> automatically is like `puts` calling `to_s`.

## Your task

`Money`, `Transaction` and `Statement` are declared in `types.go`. Give each a `String()`
method in `exercise.go`.

**Problem 1: `func (m Money) String() string`**

| Value | `String()` |
|---|---|
| `Money(1999)` | `"$19.99"` |
| `Money(5)` | `"$0.05"` |
| `Money(0)` | `"$0.00"` |
| `Money(-450)` | `"-$4.50"` |

**Problem 2: `Transaction` and `Statement`**

- `func (t Transaction) String() string` returns the date, a space, the description,
  a space, and the amount using your `Money` stringer:
  `Transaction{Date: "2024-03-01", Description: "Coffee", Amount: -450}` →
  `"2024-03-01 Coffee -$4.50"`.
  (Tip: `fmt.Sprintf("%s %s %s", ..., t.Amount)`: `%s` on a `Money` calls its `String`.)
- `func (s Statement) String() string` returns a header line followed by one line per
  transaction, each line ending in `"\n"`:

  ```
  Statement for Ana (2 transactions)
  2024-03-01 Salary $3000.00
  2024-03-02 Coffee -$4.50
  ```

  Use `"transaction"` (singular) when there is exactly 1. With no transactions it's just
  `"Statement for Ana (0 transactions)\n"`.

The tests check the output of `fmt.Sprint`, `fmt.Sprintf("%v", ...)` and `%s`, so your
`String` methods must be picked up by `fmt` automatically.

Run: `go test -v ./concepts/stringers/`
