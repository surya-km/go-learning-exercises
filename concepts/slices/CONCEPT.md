# Slices

## In plain English

A **slice** is an ordered list of values of the same type that can grow, like a Ruby
`Array`. Under the hood, a slice is a small "window" onto a hidden array. Copying a slice
copies the window, **not** the values, so two slices can end up sharing the same data.

## Creating and growing

```go
amounts := []int{100, 250, 75}  // a slice literal
amounts = append(amounts, 40)   // append returns the (possibly new) slice: always re-assign it!
len(amounts)                    // 4
amounts[0]                      // 100
amounts[1] = 300                // change an element

var empty []int                 // a nil slice: len 0, and append still works on it
none := []int{}                 // an empty but non-nil slice
buf := make([]int, 0, 10)       // length 0, room for 10 before append needs to grow
```

## Slicing: `s[low:high]`

```go
amounts[1:3]   // elements at index 1 and 2 (high is NOT included)
amounts[:2]    // the first two
amounts[2:]    // from index 2 to the end
```

Going past the end (`amounts[10]` or `amounts[:10]` on a shorter slice) **panics**, so
check `len` first.

## Sharing: the gotcha

```go
week := []int{1, 2, 3, 4, 5}
firstTwo := week[:2]
firstTwo[0] = 99
fmt.Println(week[0])            // 99, because both slices look at the same array!
```

When you need an independent copy, make a new slice and copy into it:

```go
clone := make([]int, len(week))
copy(clone, week)               // or: clone := append([]int{}, week...)
```

> **Coming from Ruby?** Go has no `.map`, `.select`, `.sum` or `.each_slice` built in.
> You write a loop. It's more typing, but it's always obvious what the code does.
> (The `slices` package has some helpers, like `slices.Contains`, but try it by hand here.)

## Your task

Implement in `exercise.go`:

**Problem 1: totals and filters**

- **`Sum(amounts []int) int`**: add up all the amounts. An empty or nil slice sums to `0`.
- **`Above(amounts []int, min int) []int`**: return a **new** slice with only the amounts
  strictly greater than `min`, in their original order. Return an empty, non-nil slice
  when nothing matches. Don't modify the input.

**Problem 2: `Pages(items []string, pageSize int) [][]string`**

Split a list of transaction IDs into pages of `pageSize` items, the way a statement shows
them 10 at a time. The last page may be shorter.

| Call | Result |
|---|---|
| `Pages([]string{"a","b","c","d","e"}, 2)` | `[[a b] [c d] [e]]` |
| `Pages([]string{"a","b"}, 5)` | `[[a b]]` |
| `Pages(nil, 3)` or `Pages(items, 0)` | `[]` (empty, non-nil) |

Each page must be an **independent copy**: changing an item in a page must not change
the original `items` slice (see "Sharing: the gotcha" above). `[][]string` is just a
slice whose elements are `[]string` slices.

Run: `go test -v ./concepts/slices/`
