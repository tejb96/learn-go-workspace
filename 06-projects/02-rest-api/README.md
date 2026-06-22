# REST API Project

## What you'll learn

- JSON CRUD with `net/http` only (no Gin/Echo)
- Middleware for logging and auth
- In-memory store with concurrency safety

## Project brief

Build `NewServer() http.Handler` with:

| Route | Auth | Behavior |
|-------|------|----------|
| `GET /health` | No | 200 OK |
| `GET /items` | Yes | List all items |
| `POST /items` | Yes | Create `{"name":"..."}` → 201 + item with `id` |
| `GET /items/{id}` | Yes | Get one or 404 |
| `PUT /items/{id}` | Yes | Update name |
| `DELETE /items/{id}` | Yes | 204 No Content |

**Auth:** require header `X-API-Key: dev-key` (except `/health`). Otherwise **401**.

**Logging middleware:** set response header **`X-Request-Logged: true`** on every request.

Use Go 1.22+ route patterns on `ServeMux` if available, or parse paths manually.

## Common mistakes

- **No mutex on map:** CRUD tests may race under `-race`.
- **Forgetting Content-Type on JSON responses:** Set `application/json`.
- **Wrong status codes:** POST → 201, DELETE → 204, missing → 404.

## Further reading

- [net/http routing (Go 1.22)](https://go.dev/blog/routing-enhancements)
- [Package net/http](https://pkg.go.dev/net/http)

## Before moving on

- [ ] `go test -v` passes
- [ ] I tested with `httptest` mentally mapped to real HTTP
- [ ] Middleware order makes sense (logging wraps auth wraps mux)
