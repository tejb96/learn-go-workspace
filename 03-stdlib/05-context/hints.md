Try for at least 30 minutes before reading this.

## Hint 1

`context.WithTimeout` returns a context that auto-cancels after a duration. Attach it to HTTP requests with `http.NewRequestWithContext`. When time runs out, the client returns an error — check with `errors.Is(err, context.DeadlineExceeded)`.

---

## Hint 2

```go
ctx, cancel := context.WithTimeout(context.Background(), d)
defer cancel()
req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
```

Read the full body with `io.ReadAll(resp.Body)`. Wrap errors with `%w` if you add context so `errors.Is` still works.

---

## Hint 3

```go
func FetchWithTimeout(url string, d time.Duration) ([]byte, error) {
    ctx, cancel := context.WithTimeout(context.Background(), d)
    defer cancel()
    return FetchWithContext(ctx, url)
}

func FetchWithContext(ctx context.Context, url string) ([]byte, error) {
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        return nil, err
    }
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return io.ReadAll(resp.Body)
}
```

This is a real timeout — not a toy sleep in your code. The server in tests deliberately stalls.
