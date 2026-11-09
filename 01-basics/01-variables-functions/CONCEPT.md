# Variables & Functions

## In plain English

Go is **strict but simple**. Every value has a fixed type that is checked before the
program runs. If you declare a variable you don't use, or import a package you don't
need, the code won't compile. That sounds annoying, but it catches a lot of bugs early.

## Variables

```go
var name string = "Ada" // explicit type
age := 30               // ":=" declares AND infers the type (most common inside functions)
age = 31                // "=" re-assigns an existing variable
// age = "thirty"       // compile error: age is an int forever
```

Variables you declare without a value get a **zero value**: `0`, `""`, `false`, or `nil`.

> **Coming from Ruby?** In Ruby, `age = "thirty"` silently changes the type.
> In Go, types never change, so the compiler catches that mistake for you.

## Functions

```go
func add(a int, b int) int {
    return a + b
}

// Parameters of the same type can share it:
func add(a, b int) int { return a + b }
```

### Multiple return values

Go functions can return more than one value. This is used *everywhere*, especially for
returning a result plus an error (you'll see that in the errors exercise).

```go
func divmod(a, b int) (int, int) {
    return a / b, a % b
}

q, r := divmod(7, 2) // q = 3, r = 1
_, r = divmod(9, 4)  // "_" ignores a value you don't need
```

### Integer division

Dividing two `int`s throws away the remainder: `7 / 2 == 3`. Use `%` to get it: `7 % 2 == 1`.

## `if` and string building

```go
if name == "" {
    name = "friend"
}

msg := "Hello, " + name + "!"          // concatenation
msg := fmt.Sprintf("Hello, %s!", name) // formatted (needs: import "fmt")
```

No parentheses around the condition, and the braces are required.

## Your task

Open `exercise.go` and implement:

1. **`Greeting(name string) string`**: return `"Hello, <name>!"`. If `name` is empty,
   use `"friend"` instead.
2. **`SplitBill(totalCents, people int) (each, leftover int)`**: split a bill evenly.
   Return how many cents each person pays and how many cents are left over.
   If `people` is 0 or less, return `0, totalCents` (nobody to split with).

> Why cents? Money should almost never be stored as a floating-point number,
> because `0.1 + 0.2 != 0.3` in floating point. Integer cents are exact.

Run: `go test -v ./01-basics/01-variables-functions/`
