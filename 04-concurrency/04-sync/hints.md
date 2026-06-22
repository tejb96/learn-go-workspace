Try for at least 30 minutes before reading this.

## Hint 1

`sync.WaitGroup` waits for goroutines. `sync.Once` runs a function exactly once. Limit workers with a buffered channel semaphore: `sem <- struct{}{}` before work, `<-sem` after.

---

## Hint 2

Preallocate `out := make([]int, len(in))`. For each index i, acquire semaphore, spawn goroutine that sets `out[i] = fn(in[i])`, release semaphore, WaitGroup tracks completion.

---

## Hint 3

```go
func ParallelMap(in []int, fn func(int) int, workers int) []int {
    if len(in) == 0 {
        return []int{}
    }
    if workers < 1 {
        workers = 1
    }
    out := make([]int, len(in))
    sem := make(chan struct{}, workers)
    var wg sync.WaitGroup
    for i, v := range in {
        wg.Add(1)
        sem <- struct{}{}
        go func(i, v int) {
            defer wg.Done()
            defer func() { <-sem }()
            out[i] = fn(v)
        }(i, v)
    }
    wg.Wait()
    return out
}
```

`OnceValue` wraps `sync.Once` and a cached int.
