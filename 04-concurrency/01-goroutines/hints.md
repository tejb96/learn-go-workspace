Try for at least 30 minutes before reading this.

## Hint 1

A goroutine is `go func() { ... }()`. Without synchronization, multiple goroutines writing `c.n++` race. `sync.Mutex` Lock/Unlock around read-modify-write fixes SafeCounter. Run `go test -race ./...` from this directory.

---

## Hint 2

```go
type SafeCounter struct {
    mu sync.Mutex
    n  int
}

func (c *SafeCounter) Inc() {
    c.mu.Lock()
    c.n++
    c.mu.Unlock()
}
```

`RunAll` uses `sync.WaitGroup`: Add(len), each goroutine Done(), Wait() at end.

---

## Hint 3

```go
func RunAll(fns []func()) {
    var wg sync.WaitGroup
    wg.Add(len(fns))
    for _, fn := range fns {
        go func(f func()) {
            defer wg.Done()
            f()
        }(fn)
    }
    wg.Wait()
}
```

Leave `UnsafeCounter` unchanged to compare with `SafeCounter` and observe `-race` output.
