# Packages

## In plain English

A **package** is a folder of Go files that belong together, like a drawer in a filing
cabinet. Every Go file starts with `package <name>`, and all files in one folder share
that name. To use code from another package you **import** it, and you can only use the
things that package has chosen to make **public**.

## Public or private? The capital letter decides

Go has no `public` / `private` keywords. The **first letter of the name** decides:

```go
package money

type Cents int                  // Capital C: exported (public), usable as money.Cents
func Format(c Cents) string {}  // exported: other packages can call money.Format
func twoDigits(n int) string {} // lowercase: unexported (private to package money)
```

This applies to everything: functions, types, constants, variables and struct fields.

## Importing a package from this module

The `go.mod` file at the root of this repo says `module go-learning-exercises`. A
package's **import path** is the module name plus its folder:

```go
package packages

import "go-learning-exercises/concepts/37-packages/money"

func Example() string {
    return money.Format(450)    // use it as <package name>.<ExportedName>
    // money.twoDigits(5)       // compile error: twoDigits is unexported
}
```

Standard library packages are imported by short paths: `import "strconv"`.
Importing a package you don't use is a compile error.

## Named types and conversions

`type Cents int` makes a new type that is stored as an `int`. Go won't silently mix
`Cents` with plain `int`, so convert explicitly when you need to:

```go
var price money.Cents = 450     // an untyped number like 450 fits either type
n := int(price)                 // Cents -> int
price = money.Cents(n + 50)     // int -> Cents
```

Two `Cents` values can be added, compared and so on, just like ints.

> **Coming from Ruby?** A Go package is a bit like a Ruby module that's also a file
> namespace. `require "money"` + `Money.format(450)` becomes `import ".../money"` +
> `money.Format(450)`. Instead of `private`, you use a lowercase name.

## Layout of this exercise

```
concepts/37-packages/
├── types.go          Item type (don't edit)
├── exercise.go       Problem 2: imports and uses the money package
└── money/
    ├── cents.go      Cents type + unexported twoDigits helper (don't edit)
    └── exercise.go   Problem 1: the money package itself
```

## Your task

**Problem 1 (in `money/exercise.go`): `Format(c Cents) string`**

Return the amount as a dollar string. Use the unexported `twoDigits` helper from
`money/cents.go` (you can call it because you're inside package `money`), and
`strconv.Itoa(n)` to turn an `int` into a string.

| Input | Result |
|---|---|
| `1234` | `"$12.34"` |
| `5` | `"$0.05"` |
| `700` | `"$7.00"` |
| `123456` | `"$1234.56"` (no thousands separators) |
| `-1250` | `"-$12.50"` (minus sign before the `$`) |

**Problem 2 (in `exercise.go`): use the money package from outside**

- **`ReceiptLine(item Item) string`**: return `"<name>: <price>"` using `money.Format`,
  e.g. `ReceiptLine(Item{Name: "Coffee", Price: 450})` → `"Coffee: $4.50"`.
- **`Total(items []Item) money.Cents`**: add up all the prices. No items → `0`.

Do Problem 1 first: `ReceiptLine` depends on `money.Format` working.

Run: `go test -v ./concepts/37-packages/...`
