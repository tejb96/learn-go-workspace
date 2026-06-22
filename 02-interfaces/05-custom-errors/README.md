# Custom Error Types

## What you'll learn

- Struct types that implement `error`
- `Unwrap()` for error chains
- Combining custom types with `errors.Is` and `errors.As`

## Concept

### Plain English

The `error` interface is tiny: one method `Error() string`. Any type with that method is an error. Struct errors carry structured data — operation name, HTTP status, field name — inspectable with `errors.As`.

Go 1.13+ error chains use **`Unwrap() error`**. Wrapping with `fmt.Errorf("...: %w", err)` and custom `Unwrap()` both let `errors.Is` and `errors.As` walk the chain.

This lesson connects **interfaces** (error is an interface) with **errors** (foundation module).

### Go syntax

```go
type OpError struct {
    Op  string
    Err error
}

func (e *OpError) Error() string {
    return e.Op + ": " + e.Err.Error()
}

func (e *OpError) Unwrap() error {
    return e.Err
}
```

```go
var ErrRetryable = errors.New("retryable")

type RetryableError struct {
    Reason string
}

func (e *RetryableError) Error() string {
    return "retryable: " + e.Reason
}

func (e *RetryableError) Unwrap() error {
    return ErrRetryable
}
```

## Annotated examples

```go
// WHY Unwrap: errors.Is(err, ErrRetryable) walks the chain without string parsing.
func (e *RetryableError) Unwrap() error {
    return ErrRetryable
}
```

```go
// WHY return nil from constructor: idiom — no error means return nil interface.
func NewOpError(op string, err error) error {
    if err == nil {
        return nil
    }
    return &OpError{Op: op, Err: err}
}
```

## Common mistakes

- **Forgetting Unwrap:** `errors.Is` won't find wrapped sentinels.
- **Error() string too vague:** `"something failed"` — loses Op/Reason data for logs.
- **Returning non-nil error interface with nil pointer:** avoid; return explicit `nil`.
- **Using custom type when sentinel suffices:** simple `var ErrNotFound = errors.New(...)` is enough for one case.

## Further reading

- [Package errors](https://pkg.go.dev/errors)
- [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)

## API spec

| Symbol | Behavior |
|--------|----------|
| `OpError` | `Error()` → `"op: inner"`; `Unwrap()` → inner |
| `NewOpError(op, err)` | nil if err nil |
| `OpFromError(err)` | Op from chain or `""` |
| `RetryableError` | `Unwrap()` → `ErrRetryable`; `Error()` includes reason |
| `MarkRetryable(err, reason)` | Wrap err as retryable |

## Before moving on

- [ ] I implement `Error()` and `Unwrap()` on custom errors
- [ ] I use `errors.As` to extract structured fields
- [ ] All tests pass: `go test -v`

## Next module

When all interface lessons pass, continue to **03-stdlib/01-strings-fmt**.
