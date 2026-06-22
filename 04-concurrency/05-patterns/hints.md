Try for at least 30 minutes before reading this.

## Hint 1

Four patterns: **worker pool** (N workers read jobs channel), **pipeline** (stage goroutine transforms stream), **fan-in** (merge channels), **done channel** (signal stop). Close output channels only after producers finish.

---

## Hint 2

Worker pool skeleton:

```go
results := make(chan int)
var wg sync.WaitGroup
for w := 0; w < workers; w++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        for j := range jobs {
            results <- j * 2
        }
    }()
}
go func() { wg.Wait(); close(results) }()
return results
```

---

## Hint 3

`ProcessUntilDone` loop:

```go
for {
    select {
    case <-ctx.Done():
        return collected
    case <-done:
        return collected
    case j, ok := <-jobs:
        if !ok { return collected }
        collected = append(collected, processor(j))
    }
}
```

Pipeline: one goroutine `for v := range in { out <- v*v }`, close out when in closes. FanIn reuses Merge pattern from select lesson.
