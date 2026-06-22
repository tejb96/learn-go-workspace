# Error Handling

## What you'll learn

- Creating errors with `errors.New`
- Wrapping errors with `fmt.Errorf` and `%w`
- Checking errors with `errors.Is` and `errors.As`
- Sentinel errors vs custom error types

## Concept

### Plain English

In Go, errors are values — usually the last return value of type `error`. There is no try/catch. Callers decide what to do: retry, log, return upstream, or fail the request.

You can define **sentinel** errors (shared variables like `ErrNotFound`) and compare them with `errors.Is`, even through wrapped layers.

**Wrapping** adds context (`"lookup user: not found"`) while preserving the original error for `errors.Is` / `errors.As` when you use `%w`.

**Custom types** (structs with an `Error() string` method) carry structured data — field names, codes — inspectable via `errors.As`.

### Go syntax

```go
var ErrNotFound = errors.New("not found")

func Lookup(id string) (string, error) {
    if id == "" {
        return "", fmt.Errorf("id required")
    }
    // ...
    return "", ErrNotFound
}

func LookupWrapped(id string) (string, error) {
    _, err := Lookup(id)
    if err != nil {
        return "", fmt.Errorf("lookup %q: %w", id, err)
    }
    return "ok", nil
}

// Caller:
if errors.Is(err, ErrNotFound) { /* handle missing */ }

var ve *ValidationError
if errors.As(err, &ve) { /* use ve.Field */ }
```

## Annotated examples

```go
// WHY %w not %v: %v stringifies and breaks the unwrap chain.
// %w keeps ErrNotFound discoverable via errors.Is.
return fmt.Errorf("load config: %w", err)
```

```go
// WHY errors.Is instead of err == ErrNotFound:
// wrapped errors are not equal with == but Is walks the chain.
if errors.Is(err, ErrNotFound) {
    return http.StatusNotFound
}
```

```go
// WHY custom type: sentinel strings cannot carry Field and Code.
type ValidationError struct {
    Field string
    Code  string
}
func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation: %s %s", e.Field, e.Code)
}
```

## Common mistakes

- **Using `==` on wrapped errors:** Use `errors.Is` / `errors.As`.
- **Wrapping with `%v`:** Loses unwrap chain; use `%w`.
- **Returning nil error as non-nil interface:** `var err *MyError = nil; return err` can be non-nil interface; return explicit `nil`.
- **String matching:** `strings.Contains(err.Error(), "not found")` breaks when messages change.

## Further reading

- [Package errors](https://pkg.go.dev/errors)
- [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)

## API spec

Package name: `errorslesson` (directory `03-errors`).

| Symbol | Behavior |
|--------|----------|
| `ErrNotFound` | Sentinel, message `"not found"` |
| `ValidationError` | Fields `Field`, `Code`; `Error()` returns `"validation failed"` |
| `Wrap(err, msg)` | `fmt.Errorf("%s: %w", msg, err)`; result message for `ErrNotFound` + `"lookup user"` is `"lookup user: not found"` |
| `Find(store, id)` | Empty `id` → `*ValidationError` with `id`/`required`; missing key → `ErrNotFound`; else value |

## Before moving on

- [ ] I can create and return sentinel errors
- [ ] I wrap with `%w` and check with `errors.Is`
- [ ] I use `errors.As` for custom error types
- [ ] All tests pass: `go test -v`
