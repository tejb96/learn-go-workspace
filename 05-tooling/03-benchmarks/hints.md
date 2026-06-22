Try for at least 30 minutes before reading this.

## Hint 1

Benchmarks use `func BenchmarkX(b *testing.B)` and run with `go test -bench=.`. `b.N` is how many iterations the framework chose. Fix unit tests first with a simple recursive or iterative Fib.

---

## Hint 2

Iterative Fib avoids exponential recursion:

```go
func Fib(n int) int {
    if n < 2 {
        return n
    }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}
```

---

## Hint 3

`FibCached`: if v, ok := cache[n]; ok { return v }. Compute with Fib or inner loop, store in cache, return. Run benchmarks:

```bash
go test -bench=. -benchmem
```

Compare `BenchmarkFib` vs `BenchmarkFibCached` after implementing cache (cached may win on repeated same n).
