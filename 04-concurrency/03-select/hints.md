Try for at least 30 minutes before reading this.

## Hint 1

`select` waits on multiple channel operations; the first ready case runs. Use `default` for non-blocking attempts. Closed channels yield zero values immediately with ok=false on receive.

---

## Hint 2

`FirstReady`: loop with select over cases `case v := <-ch: return v, nil`. Track how many channels are closed; when all closed, return `ErrNoReadyChannel`. Do not use `default` if you need to wait.

---

## Hint 3

For `Merge`, start one goroutine per input channel copying to shared output, use WaitGroup, close output when all inputs done:

```go
out := make(chan int)
var wg sync.WaitGroup
for _, ch := range chs {
    wg.Add(1)
    go func(c <-chan int) {
        defer wg.Done()
        for v := range c {
            out <- v
        }
    }(ch)
}
go func() { wg.Wait(); close(out) }()
return out
```
