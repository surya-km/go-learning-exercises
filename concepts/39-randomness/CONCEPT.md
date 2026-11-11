# Randomness

## In plain English

The `math/rand/v2` package gives you random-looking numbers for things like games,
shuffling a playlist, or picking a sample of transactions to review. They come from a
**generator**: a formula that starts from a **seed** number. The same seed always
produces the same sequence. That sounds like a flaw, but it's exactly what makes random
code testable.

```go
import "math/rand/v2"

rand.IntN(6)        // 0, 1, 2, 3, 4 or 5   (top-level functions use a random seed)
rand.IntN(6) + 1    // a dice roll: 1 to 6
rand.Float64()      // somewhere in [0.0, 1.0)
```

## Your own generator with a fixed seed

Instead of the top-level functions, functions can **accept** a generator. In real code
you pass one with a random seed. In tests you pass one with a fixed seed, so every run
behaves the same:

```go
r := rand.New(rand.NewPCG(1, 2))   // a generator seeded with 1 and 2

r.IntN(10)                          // a number from 0 to 9
r.Shuffle(len(s), func(i, j int) {  // shuffle a slice in place
    s[i], s[j] = s[j], s[i]
})
r.Perm(5)                           // a random ordering of 0..4, e.g. [3 0 4 1 2]
```

Functions that take `r *rand.Rand` don't care which seed it has. The `*` means it's a
pointer to the generator, so calls share and advance the same generator. Pointers are
covered later. For now, just pass `r` along.

**Warning:** `math/rand/v2` is **not** safe for passwords, one-time passcodes in real
banking apps, or anything security related, because an attacker could predict it. Use
`crypto/rand` for those. This exercise uses `math/rand/v2` because it's for practice.

> **Coming from Ruby?** `rand(6)` is `rand.IntN(6)`, `arr.shuffle` is `r.Shuffle`, and
> `Random.new(42)` is `rand.New(rand.NewPCG(42, 0))`.

## Copying a slice

Problem 2 must not change the caller's slice, so you'll need a copy first:

```go
copied := make([]string, len(items))
copy(copied, items)       // or: copied := slices.Clone(items)
```

## Your task

**Problem 1:** implement **`VerificationCode(r *rand.Rand, length int) string`**.

Return a string of `length` random digits (`'0'` to `'9'`), like `"048213"`. Leading
zeros are allowed. If `length` is 0 or less, return `""`. Use `r` for all randomness, not
the top-level `rand` functions.

**Problem 2:** implement **`ShuffledCopy(r *rand.Rand, items []string) []string`**.

Return a **new** slice with the same items in a random order (for example, picking the
order in which to review flagged transactions). The original `items` slice must not be
changed. For an empty or `nil` input, return an empty slice or `nil`.

The tests can't check for exact values, so they check properties instead: the length, that
only digits are used, that every digit shows up across many codes, that the same seed
gives the same result, and that no items are lost or added by the shuffle.

Run: `go test -v ./concepts/39-randomness/`
