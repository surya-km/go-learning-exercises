# Range Iteration

## In plain English

`for ... range` walks through a collection one item at a time and hands you each item
in turn. The same keyword works on slices, maps, strings and (since Go 1.22) plain
integers. What it hands you depends on what you range over.

## Slices: index and value

```go
amounts := []int{100, 250, 75}
for i, amt := range amounts {   // i = 0, 1, 2   amt = 100, 250, 75
    fmt.Println(i, amt)
}
for _, amt := range amounts { } // only the value ("_" ignores the index)
for i := range amounts { }      // only the index
```

`amt` is a **copy**. Changing it doesn't change the slice. To change the slice, write
`amounts[i] = ...`.

## Maps: key and value, in random order

```go
budget := map[string]int{"food": 40000, "rent": 120000}
for category, cents := range budget {
    fmt.Println(category, cents)
}
```

⚠️ Go **shuffles map order on purpose**, so it can be different on every run. Never rely
on it. If two keys could tie, decide the winner with an explicit rule.

## Strings: byte position and rune (character)

```go
for i, r := range "café" {
    fmt.Println(i, string(r))   // 0 c, 1 a, 2 f, 3 é
}
```

`r` is a `rune` (a Unicode character), not a byte, so accented letters and emoji come
out whole. Compare runes with single quotes: `r == ' '`, `r >= '0' && r <= '9'`.
Build a string up with `strings.Builder`:

```go
var b strings.Builder
b.WriteRune('*')
b.WriteRune(r)
b.String()                      // "*c"
```

## Integers (Go 1.22+): count from 0 up to n-1

```go
for i := range 3 {              // i = 0, 1, 2
    fmt.Println(i)
}
for range 3 { }                 // just repeat 3 times
```

> **Coming from Ruby?** `range` covers `.each`, `.each_with_index`, `hash.each`,
> `string.each_char` and `3.times`, all with one keyword.

## Your task

Implement in `exercise.go`:

**Problem 1: slices and integers**

- **`RunningBalance(startCents int, amounts []int) []int`** (range over a slice):
  apply each amount to the balance in order and return the balance after each one.
  `RunningBalance(1000, []int{-200, 500, -50})` → `[800, 1300, 1250]`. With no amounts,
  return an empty, non-nil slice.
- **`InstallmentPlan(totalCents, months int) []int`** (range over an integer): split a
  purchase into `months` payments that differ by at most one cent. Any leftover cents go
  on the **earliest** payments. `InstallmentPlan(1000, 3)` → `[334, 333, 333]`. If
  `months` is 0 or less, return an empty, non-nil slice.

**Problem 2: strings and maps**

- **`MaskCard(card string) string`** (range over a string): replace every digit with
  `*` except the **last four digits**. Keep all other characters (spaces, dashes) as they
  are. `MaskCard("4111 1111 1111 1234")` → `"**** **** **** 1234"`. If there are four
  digits or fewer, nothing is masked.
- **`BiggestCategory(totals map[string]int) (string, int)`** (range over a map): return
  the category with the largest total and that total. If two categories tie, return the
  one that comes **first alphabetically** (`"food" < "rent"` works on strings). For an
  empty map, return `"", 0`.

Run: `go test -v ./concepts/range-iteration/`
