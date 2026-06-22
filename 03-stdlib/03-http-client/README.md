# HTTP Client

## What you'll learn

- Making GET requests with `net/http`
- Decoding JSON responses
- Passing `context.Context` into requests
- Testing HTTP clients with `httptest.Server`

## Concept

### Plain English

Go's **`net/http`** client sends requests and reads responses. You create a request (optionally with context for cancellation), call `Client.Do`, check the status code, read the body, and decode JSON if needed.

This lesson covers the **client only**. Building a server is deferred to the REST API project in Module 6 — learn to call APIs first.

### Go syntax

```go
req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
resp, err := http.DefaultClient.Do(req)
if err != nil {
    return err
}
defer resp.Body.Close()

if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("status: %s", resp.Status)
}
return json.NewDecoder(resp.Body).Decode(&dest)
```

## Annotated examples

```go
// WHY context on request: timeouts and cancellation propagate to the transport.
req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
```

```go
// WHY httptest in tests: no real network; full control over responses.
srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(payload)
}))
defer srv.Close()
clientURL := srv.URL
```

## Common mistakes

- **Forgetting `defer resp.Body.Close()`:** Leaks connections.
- **Ignoring status codes:** 404 body is often HTML, not JSON.
- **Using `http.Get` without context:** Prefer `NewRequestWithContext`.
- **Not draining/closing body on errors:** Still close the body on non-2xx.

## Further reading

- [Package net/http](https://pkg.go.dev/net/http)
- [Package net/http/httptest](https://pkg.go.dev/net/http/httptest)

## API spec

| Function | Behavior |
|----------|----------|
| `FetchJSON(ctx, url, dest)` | GET + JSON decode; error on non-2xx or bad JSON |
| `GetStatus(ctx, url)` | GET; return status code |

## Before moving on

- [ ] I close response bodies
- [ ] I check HTTP status before decoding
- [ ] I know server building comes later
- [ ] All tests pass: `go test -v`
