Try for at least 30 minutes before reading this.

## Hint 1

This lesson is **HTTP client only** — no server yet. Use `http.NewRequestWithContext`, `http.DefaultClient.Do`, read the body, and `json.NewDecoder`. Always close `resp.Body`.

---

## Hint 2

Check `resp.StatusCode` before decoding — return `fmt.Errorf` for status >= 300. Pass `ctx` into the request so cancellation works in later lessons. `GetStatus` can share the same request pattern but skip JSON.

---

## Hint 3

```go
func FetchJSON(ctx context.Context, url string, dest any) error {
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        return err
    }
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        return fmt.Errorf("unexpected status: %s", resp.Status)
    }
    return json.NewDecoder(resp.Body).Decode(dest)
}
```

Building a server comes in Module 6 projects.
