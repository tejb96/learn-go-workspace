Try for at least 30 minutes before reading this.

## Hint 1

Unbuffered channels block until sender and receiver rendezvous. Buffered channels block only when full (send) or empty (receive). ASCII in README: `A ---[ ch ]--- B`.

---

## Hint 2

`PingPong`: one unbuffered `chan struct{}`, goroutine A sends then receives n times, B receives then sends n times. Use `select` with `time.After` only in `SendWithTimeout`, not PingPong.

---

## Hint 3

```go
func SendWithTimeout(ch chan<- int, v int, d time.Duration) bool {
    select {
    case ch <- v:
        return true
    case <-time.After(d):
        return false
    }
}
```

For `BufferedCollect`, loop `count` times: `v := <-ch`, append to slice.
