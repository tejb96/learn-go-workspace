# Testing

## What you'll learn

- Table-driven tests with `t.Run` subtests
- `t.Helper()` for test helpers
- `testing.T` vs `testing.B` (benchmarks come next lesson)

## Concept

### Plain English

Go tests live in `*_test.go` files in the same package (or `_test` suffix package for black-box tests). **`testing.T`** is passed to each test function for failures, logging, and subtests.

**Subtests** (`t.Run`) group related cases — run one: `go test -run TestFoo/bar`.

**`t.Helper()`** marks a helper function so when it calls `t.Fatal`, the failure line points to the test caller, not the helper. See `helper_test.go` in this lesson (do not edit it).

**`testing.B`** is for benchmarks (`func BenchmarkX(b *testing.B)`), not unit tests. Same `go test` command with `-bench`.

### Go syntax

```go
func TestThing(t *testing.T) {
    tests := []struct {
        name string
        in   int
        want int
    }{
        {name: "zero", in: 0, want: 0},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Thing(tt.in)
            if got != tt.want {
                t.Fatalf("got %d want %d", got, tt.want)
            }
        })
    }
}
```

```go
func helper(t *testing.T, got, want int) {
    t.Helper()
    if got != want {
        t.Fatalf("got %d want %d", got, want)
    }
}
```

## Annotated examples

```go
// WHY t.Run: filter failures, parallel subtests (t.Parallel), clearer output.
t.Run("case name", func(t *testing.T) { ... })
```

```go
// WHY t.Helper: stack traces show which subtest case failed, not assertEqual line.
t.Helper()
```

## Common mistakes

- **Testing in `main`:** Use `go test`.
- **Forgetting t.Helper in shared asserts:** Harder to debug failures.
- **Parallel tests sharing mutable state:** Race conditions — use `-race`.
- **Confusing T and B:** Benchmarks use `b *testing.B` and `b.N`.

## Further reading

- [Package testing](https://pkg.go.dev/testing)
- [Go blog — Subtests](https://go.dev/blog/subtests)

## API spec

| Function | Behavior |
|----------|----------|
| `IsPalindrome(s)` | Empty → true; case-sensitive |
| `SumPositive(nums)` | Sum of values > 0 |

## Before moving on

- [ ] I ran a single subtest with `-run`
- [ ] I understand why `helper_test.go` uses `t.Helper()`
- [ ] I know `testing.B` is for benchmarks
- [ ] All tests pass: `go test -v`
