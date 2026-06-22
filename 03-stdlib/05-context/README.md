# context

## What you'll learn

- `context.WithTimeout` and cancellation
- Propagating context through HTTP requests
- Detecting `DeadlineExceeded` and `Canceled`

## Concept

### Plain English

**Context** carries deadlines, cancellation signals, and request-scoped values across API boundaries. When a caller gives up (user navigates away, timeout fires), context cancellation tells downstream work to stop — close connections, abort slow queries, skip wasted CPU.

For HTTP clients, attach context to the request. If the server takes too long, the client aborts and returns `context.DeadlineExceeded`. This lesson uses a **real slow server** in tests, not a fake sleep in your implementation pretending to timeout.

### Go syntax

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
resp, err := http.DefaultClient.Do(req)
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        // timed out
    }
    return err
}
defer resp.Body.Close()
```

## Annotated examples

```go
// WHY defer cancel(): releases timer resources even if request succeeds early.
ctx, cancel := context.WithTimeout(context.Background(), d)
defer cancel()
```

```go
// WHY pass ctx to request, not just Client.Timeout:
// consistent pattern for RPC, DB, and HTTP; supports manual Cancel too.
req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
```

## Common mistakes

- **Storing context in a struct:** Pass as first function parameter instead.
- **Using context.Background() everywhere:** Use timeout/cancel at the edge (main, handler).
- **Ignoring cancel func:** Always defer cancel to free timer.
- **Checking error strings:** Use `errors.Is(err, context.DeadlineExceeded)`.

## Further reading

- [Package context](https://pkg.go.dev/context)
- [Go blog — context](https://go.dev/blog/context)

## API spec

| Function | Behavior |
|----------|----------|
| `FetchWithTimeout(url, d)` | GET with timeout; `DeadlineExceeded` on slow response |
| `FetchWithContext(ctx, url)` | GET honoring ctx; returns body bytes |

## Before moving on

- [ ] I attach context to outbound HTTP requests
- [ ] I use `errors.Is` for deadline/cancel errors
- [ ] All tests pass: `go test -v`

## Next module

When all stdlib lessons pass, continue to **04-concurrency/01-goroutines**.
