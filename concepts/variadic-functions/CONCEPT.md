# Variadic Functions

## In plain English

A **variadic** function accepts *any number* of arguments of the same type: zero, one,
or a hundred. You've already used one: `fmt.Println("a", "b", "c")`. Inside the
function, all those extra arguments arrive packed together as a single slice.

```go
func Sum(amounts ...int) int {   // "...int" means "zero or more ints"
    total := 0
    for _, a := range amounts {  // amounts is a []int inside the function
        total += a
    }
    return total
}

Sum()               // 0
Sum(100)            // 100
Sum(100, 250, 75)   // 425
```

## Passing a slice you already have: `s...`

If your values are already in a slice, you can't pass the slice directly. Add `...`
after it to "unpack" it into the variadic parameter:

```go
prices := []int{100, 250, 75}
Sum(prices...)      // 425
// Sum(prices)      // compile error: cannot use []int as int
```

## Rules

- Only the **last** parameter can be variadic, and only one per function:
  ```go
  func Tag(label string, amounts ...int) // ok
  // func Bad(amounts ...int, label string) // compile error
  ```
- You can mix fixed arguments with the variadic ones: `Tag("food", 100, 200)`.
- With no extra arguments, the slice is `nil` (length 0). Your code must handle that.
- `append` is variadic too: `all = append(all, more...)` joins two slices.

> **Coming from Ruby?** This is Go's version of the splat: `def sum(*amounts)` becomes
> `func Sum(amounts ...int)`, and `sum(*prices)` becomes `Sum(prices...)`.

## Your task

Implement in `exercise.go`:

**Problem 1: `MaxAmount(amounts ...int) (int, bool)`**

Return the largest amount and `true`. If no amounts were passed, return `0, false`.

| Call | Result |
|---|---|
| `MaxAmount(300, 1200, 450)` | `1200, true` |
| `MaxAmount(-50, -10)` | `-10, true` (refunds can be negative) |
| `MaxAmount()` | `0, false` |

**Problem 2: `AddToBill(billCents int, items ...int) int`**

Start with an existing bill and add every item price to it. `billCents` is a normal
parameter, and `items` is the variadic one. With no items, return `billCents` unchanged.

| Call | Result |
|---|---|
| `AddToBill(1000, 250, 399)` | `1649` |
| `AddToBill(1000)` | `1000` |
| `AddToBill(0, []int{100, 200}...)` | `300` |

The tests call your functions both with separate arguments *and* with a slice plus `...`,
so try both in your head as you write them.

Run: `go test -v ./concepts/variadic-functions/`
