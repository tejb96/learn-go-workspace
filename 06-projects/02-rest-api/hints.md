Try for at least 30 minutes before reading this.

## Hint 1

Use `http.NewServeMux` or `ServeMux` patterns: register `/health`, `/items`, `/items/`. Wrap the mux with middleware functions: `func middleware(next http.Handler) http.Handler`.

---

## Hint 2

Auth middleware: if `r.Header.Get("X-API-Key") != "dev-key"` and path is not `/health`, return 401. Logging middleware: set response header `X-Request-Logged: true` and optionally log method/path.

Store items in `map[string]Item` with `sync.RWMutex`. Generate IDs with `fmt.Sprintf("id-%d", len+1)` or `crypto/rand`.

---

## Hint 3

```go
mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
})

mux.HandleFunc("POST /items", func(w http.ResponseWriter, r *http.Request) {
    var in struct { Name string `json:"name"` }
    json.NewDecoder(r.Body).Decode(&in)
    // create, w.WriteHeader(http.StatusCreated), json.NewEncoder(w).Encode(item)
})
```

Go 1.22+ supports method-aware patterns like `"GET /items/{id}"`. Wrap: `return loggingMiddleware(authMiddleware(mux))`.
