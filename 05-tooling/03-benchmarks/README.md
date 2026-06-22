# Benchmarks

## What you'll learn

- Writing benchmarks with `testing.B`
- Running benchmarks: `go test -bench`
- `testing.T` vs `testing.B`

## Concept

### Plain English

**Unit tests** (`testing.T`) assert correctness. **Benchmarks** (`testing.B`) measure speed and allocations. Same `go test` driver, different function signature.

The framework adjusts `b.N` until timing is stable. Your loop runs `b.N` times:

```go
for i := 0; i < b.N; i++ {
    Fib(20)
}
```

Run: `go test -bench=.` or `go test -bench=Fib -benchmem`

### testing.T vs testing.B

| | `testing.T` | `testing.B` |
|---|-------------|-------------|
| Purpose | Correctness | Performance |
| Run | `go test` | `go test -bench` |
| Loop | Your table/subtests | `for i := 0; i < b.N; i++` |
| Fail | `t.Fatal` | `b.Fatal` |

## Annotated examples

```go
// WHY ResetTimer: exclude setup (map alloc) from benchmark timing.
cache := make(map[int]int)
b.ResetTimer()
for i := 0; i < b.N; i++ {
    FibCached(20, cache)
}
```

## Common mistakes

- **Benchmarking empty loop:** Must call code under test inside loop.
- **Including init in timed section:** Use `b.ResetTimer()` after setup.
- **Comparing bench results from different machines:** Treat as relative only.
- **Optimizing before correctness:** Tests must pass first.

## Further reading

- [Package testing — benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks)

## API spec

| Function | Behavior |
|----------|----------|
| `Fib(n)` | Standard Fibonacci |
| `FibCached(n, cache)` | Same result, use cache map |

## Before moving on

- [ ] Tests pass: `go test -v`
- [ ] I ran `go test -bench=. -benchmem`
- [ ] I understand `b.N` and `testing.B`
