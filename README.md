# Go Learning Exercises

A hands-on repo for learning Go as a team. It has **39 concepts**, named after the
[Exercism Go track's concepts](https://github.com/exercism/go/tree/main/concepts). Each
concept folder has 1–2 small coding problems.

Every concept folder contains:

| File | What it is |
|---|---|
| `CONCEPT.md` | A short plain-English explanation of the idea, with examples. Read this first. |
| `exercise.go` | The code **you** edit. Functions are stubbed out with `TODO`s. |
| `exercise_test.go` | Tests that pass once your solution is correct. Don't edit these. |
| `types.go` (some folders) | Types the exercise uses. Read it, but don't edit it. |
| `.meta/solution.go` | A reference answer. Try not to peek until your tests pass! |

**The tests are the teacher.** Each exercise starts with failing tests, and you're done
when they all pass.

## Setup

1. Install Go 1.25+ with `brew install go`, and check it with `go version`.
2. Clone this repo and `cd` into it.
3. (Recommended) Install the Go extension for VS Code, or use GoLand.

## How to do an exercise

```sh
# 1. Read the concept
cat concepts/basics/CONCEPT.md

# 2. Run the tests. They fail at first, which is expected.
go test ./concepts/basics/

# 3. Edit concepts/basics/exercise.go until the tests pass
go test -v ./concepts/basics/
```

Tip: `go test -v -run TestSplitBill ./concepts/basics/` runs a single test.

> Running `go test ./...` from the root shows failures for every exercise you haven't
> solved yet. That's normal, so work on one folder at a time.

## Learning path

The folders are listed alphabetically on disk, but they're meant to be done **in this
order**. Later concepts build on earlier ones.

### 1. Foundations
| # | Concept | What you build |
|---|---|---|
| 1 | [basics](concepts/basics) | `Greeting`, `SplitBill` |
| 2 | [variables](concepts/variables) | `MonthEnd`, `Transfer` |
| 3 | [comments](concepts/comments) | Doc comments, `FormatAmount`, `Describe` |
| 4 | [constants](concepts/constants) | A `Frequency` enum with `iota`, `YearlyCostCents` |
| 5 | [zero-values](concepts/zero-values) | `ApplyDefaults`, `AddSpend` |

### 2. Numbers and logic
| # | Concept | What you build |
|---|---|---|
| 6 | [numbers](concepts/numbers) | `SplitBill` (fair split), `MakeChange` |
| 7 | [arithmetic-operators](concepts/arithmetic-operators) | `TipCents`, `RoundUp` |
| 8 | [floating-point-numbers](concepts/floating-point-numbers) | `ApplyRate`, `AlmostEqual` |
| 9 | [type-conversion](concepts/type-conversion) | `AverageCents`, `OrderSummary` |
| 10 | [booleans](concepts/booleans) | `CanWithdraw`, `NeedsReview` |
| 11 | [comparison](concepts/comparison) | `CompareAmounts`, `EarlierDate` |
| 12 | [conditionals-if](concepts/conditionals-if) | `OverdraftFee`, `BudgetStatus` |
| 13 | [conditionals-switch](concepts/conditionals-switch) | `Category`, `CreditRating` |

### 3. Functions and loops
| # | Concept | What you build |
|---|---|---|
| 14 | [functions](concepts/functions) | `Greeting`, `TipCents`, `TotalWithTip` |
| 15 | [multiple-return-values](concepts/multiple-return-values) | `SplitBill`, `CoinChange` |
| 16 | [for-loops](concepts/for-loops) | `MonthsToPayOff`, `CountWeekdays` |
| 17 | [variadic-functions](concepts/variadic-functions) | `MaxAmount`, `AddToBill` |

### 4. Collections
| # | Concept | What you build |
|---|---|---|
| 18 | [slices](concepts/slices) | `Sum`, `Above`, `Pages` |
| 19 | [maps](concepts/maps) | `TotalByCategory`, `OverBudget` |
| 20 | [range-iteration](concepts/range-iteration) | `RunningBalance`, `InstallmentPlan`, `MaskCard`, `BiggestCategory` |

### 5. Text
| # | Concept | What you build |
|---|---|---|
| 21 | [strings](concepts/strings) | `MaskAccount`, `ShortID` |
| 22 | [runes](concepts/runes) | `Truncate`, `Initials` (handles é, € and emoji) |
| 23 | [strings-package](concepts/strings-package) | `CleanMerchant`, `ParseTags` |
| 24 | [strings-builder](concepts/strings-builder) | `Receipt`, `GroupDigits` |
| 25 | [string-formatting](concepts/string-formatting) | `StatementLine`, `SpendingShare` |
| 26 | [regular-expressions](concepts/regular-expressions) | `ValidRoutingNumber`, `ExtractAmounts` |

### 6. Your own types
| # | Concept | What you build |
|---|---|---|
| 27 | [structs](concepts/structs) | `NewAccount`, `Summarize` |
| 28 | [type-definitions](concepts/type-definitions) | `Cents.Format`, `Cents.Split`, `YearlyCost` |
| 29 | [methods](concepts/methods) | `Bill.IsDueBy`, `Bill.Label`, `Bills.TotalDueBy` |
| 30 | [pointers](concepts/pointers) | `Account.Deposit`, `Account.Withdraw`, `Transfer` |
| 31 | [stringers](concepts/stringers) | `String()` for `Money`, `Transaction`, `Statement` |
| 32 | [interfaces](concepts/interfaces) | A `Payment` interface on 3 types, `MonthlyTotal`, `Biggest` |
| 33 | [type-assertion](concepts/type-assertion) | `Describe` (type switch), `LookupString`, `LookupCents` |

### 7. Errors
| # | Concept | What you build |
|---|---|---|
| 34 | [errors](concepts/errors) | `ParseAmount` with sentinel errors |
| 35 | [error-wrapping](concepts/error-wrapping) | `ProcessPayment`, `DeclineCode` (`errors.As`), `RootCause`, `ProcessBatch` (`errors.Join`) |

### 8. More tools
| # | Concept | What you build |
|---|---|---|
| 36 | [first-class-functions](concepts/first-class-functions) | `Filter`, `SpendingTracker` (closure) |
| 37 | [packages](concepts/packages) | A `money` sub-package, `ReceiptLine`, `Total` |
| 38 | [time](concepts/time) | `DaysInMonth`, `NextDueDate` (month-end safe), `DaysBetween`, `FormatDue` |
| 39 | [randomness](concepts/randomness) | `VerificationCode`, `ShuffledCopy` |

### Planned next
- `concurrency`: goroutines, channels, `sync.WaitGroup`, `context`
- `testing`: writing your own table-driven tests
- `real-world`: HTTP handlers, JSON, and the patterns our services actually use

Want to add one? See [CONTRIBUTING.md](CONTRIBUTING.md).

## Checking the reference solutions

Maintainers can confirm that every reference solution passes its tests:

```sh
./scripts/verify-solutions.sh
```

CI runs this on every push.

## More practice

- [exercism.org Go track](https://exercism.org/tracks/go): free, ordered, with optional mentoring
- [A Tour of Go](https://go.dev/tour/): the official interactive intro
- [Go by Example](https://gobyexample.com/): short annotated snippets
- [Effective Go](https://go.dev/doc/effective_go): idiomatic style
