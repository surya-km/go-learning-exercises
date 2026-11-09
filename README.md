# Go Learning Exercises

A small, hands-on repo for learning Go as a team. Each exercise has:

| File | What it is |
|---|---|
| `CONCEPT.md` | A short plain-English explanation of the idea, with examples. Read this first. |
| `exercise.go` | The code **you** edit. Functions are stubbed out with `TODO`s. |
| `exercise_test.go` | Tests that pass once your solution is correct. Don't edit these. |
| `.meta/solution.go` | A reference answer. Try not to peek until your tests pass! |

The idea is borrowed from [Exercism](https://exercism.org/tracks/go): **the tests are the teacher.**
You start with failing tests and work until they're green.

## Setup

1. Install Go 1.25+ — `brew install go` (check with `go version`).
2. Clone this repo and `cd` into it.
3. (Recommended) Install the Go extension for VS Code, or use GoLand.

## How to do an exercise

```sh
# 1. Read the concept
cat 01-basics/01-variables-functions/CONCEPT.md

# 2. Run the tests — they fail at first, that's expected
go test ./01-basics/01-variables-functions/

# 3. Edit exercise.go until the tests pass
go test -v ./01-basics/01-variables-functions/
```

Tip: `go test -run TestSplitBill ./01-basics/01-variables-functions/` runs a single test.

> Running `go test ./...` from the root will show failures for every exercise you
> haven't solved yet — that's normal. Work on one folder at a time.

## Learning path

Do them in order — later exercises assume earlier concepts.

| # | Exercise | Concepts |
|---|---|---|
| 1 | [01-basics/01-variables-functions](01-basics/01-variables-functions) | variables, `:=`, functions, multiple return values, `if` |
| 2 | [01-basics/02-slices-maps](01-basics/02-slices-maps) | slices, `append`, `range`, maps, structs (intro) |
| 3 | [02-errors/01-parse-amount](02-errors/01-parse-amount) | the `error` type, `if err != nil`, sentinel errors, wrapping with `%w`, `errors.Is` |

### Planned next

- `03-types-interfaces` — methods, interfaces, embedding
- `04-testing` — table-driven tests, writing your own tests
- `05-concurrency` — goroutines, channels, `sync.WaitGroup`, `context`
- `06-real-world` — HTTP handlers, JSON, the patterns our services actually use

Want to add one? See [CONTRIBUTING.md](CONTRIBUTING.md).

## Checking the reference solutions

Maintainers can confirm every reference solution passes its tests:

```sh
./scripts/verify-solutions.sh
```

CI runs this on every push.

## More practice

- [exercism.org Go track](https://exercism.org/tracks/go) — free, ordered, with optional mentoring
- [A Tour of Go](https://go.dev/tour/) — the official interactive intro
- [Go by Example](https://gobyexample.com/) — short annotated snippets
- [Effective Go](https://go.dev/doc/effective_go) — idiomatic style
