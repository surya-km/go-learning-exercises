# Basics

## In plain English

Every Go file starts by saying which **package** it belongs to. A package is just a folder
of related code. Inside it you write **functions**: named blocks of code that take some
inputs, do some work, and hand back a result. That's most of what a Go program is.

```go
package greeting          // this file belongs to the "greeting" package

import "fmt"              // use code from the standard library's fmt package

// Hello is exported (usable from other packages) because it starts with a capital letter.
func Hello(name string) string {
    return "Hi " + name
}

// shout is unexported: only code inside this package can call it.
func shout(s string) string { return s + "!" }
```

## Functions

```go
func add(a int, b int) int { return a + b }
func add(a, b int) int { return a + b }      // same type? share it
```

### Multiple return values

A Go function can return more than one value:

```go
func divmod(a, b int) (int, int) {
    return a / b, a % b
}

q, r := divmod(7, 2) // q = 3, r = 1
_, r = divmod(9, 4)  // "_" throws away a value you don't need
```

Dividing two `int`s drops the remainder (`7 / 2 == 3`). `%` gives you the remainder (`7 % 2 == 1`).

## Building strings

```go
"Hello, " + name + "!"            // join strings with +
fmt.Sprintf("Hello, %s!", name)   // or fill in a template: %s = string, %d = int
```

> **Coming from Ruby?** There's no implicit return. You must write `return`. And there's
> no `"#{name}"` interpolation; use `+` or `fmt.Sprintf`.

## Your task

**Problem 1.** Implement `Greeting(name string) string`.

- `Greeting("Ada")` returns `"Hello, Ada!"`.
- An empty name returns `"Hello, friend!"`.

**Problem 2.** Implement `SplitBill(totalCents, people int) (each, leftover int)`. Money is in
integer **cents** (`$12.34` is `1234`).

- `SplitBill(1000, 3)` returns `333, 1`: each person pays 333 cents and 1 cent is left over.
- `SplitBill(900, 3)` returns `300, 0`.
- If `people` is `0` or negative, return `0, totalCents` (nobody to split with, so it's all left over).

Run: `go test -v ./concepts/01-basics/`
