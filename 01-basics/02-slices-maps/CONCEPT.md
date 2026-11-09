# Slices & Maps

## In plain English

- A **slice** is an ordered list that can grow, like a Ruby `Array`.
- A **map** is a key → value lookup, like a Ruby `Hash`.
- A **struct** is a named bundle of fields. Think of it as a lightweight class with
  no methods yet. Structs are covered properly later, but you'll *use* one here.

## Slices

```go
amounts := []int{100, 250, 75}  // a slice literal
amounts = append(amounts, 40)   // append returns a NEW slice; always re-assign it!
fmt.Println(len(amounts))       // 4
fmt.Println(amounts[0])         // 100

var empty []int                 // a nil slice: len 0, and append still works on it
```

### Looping with `range`

```go
for i, amt := range amounts {   // i = index, amt = value
    fmt.Println(i, amt)
}
for _, amt := range amounts {   // ignore the index with "_"
    total += amt
}
```

> **Coming from Ruby?** Go has no `.map`, `.select` or `.sum` built in. You write a
> `for ... range` loop. It's more typing, but it's always obvious what the code does.

## Maps

```go
totals := map[string]int{}      // empty map, ready to use
totals["food"] += 500           // missing keys read as the zero value (0), so += just works
totals["rent"] = 120000

v, ok := totals["travel"]       // "comma ok": ok is false if the key isn't there
delete(totals, "food")
```

⚠️ `var m map[string]int` creates a **nil map**. Reading from it is fine, but *writing*
to it panics. Use `map[string]int{}` or `make(map[string]int)`.

## Structs (sneak peek)

```go
type Transaction struct {
    Category    string
    AmountCents int
}

t := Transaction{Category: "food", AmountCents: 1250}
fmt.Println(t.Category) // "food"
```

`Transaction` is already defined for you in `types.go`.

## Your task

Implement in `exercise.go`:

1. **`Sum(amounts []int) int`**: add up all the amounts. An empty slice sums to `0`.
2. **`Above(amounts []int, min int) []int`**: return a new slice with only the amounts
   strictly greater than `min`, in their original order.
3. **`TotalByCategory(txns []Transaction) map[string]int`**: return a map from each
   category to the total cents spent in it. With no transactions, return an empty
   (non-nil) map.

Run: `go test -v ./01-basics/02-slices-maps/`
