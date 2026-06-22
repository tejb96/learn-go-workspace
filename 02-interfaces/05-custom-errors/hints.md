Try for at least 30 minutes before reading this.

## Hint 1

Custom error types implement the `error` interface with `Error() string`. To participate in `errors.Is` and `errors.As` unwrap chains, also implement `Unwrap() error` returning the wrapped error.

---

## Hint 2

`OpError.Error()` should return `e.Op + ": " + e.Err.Error()`. `Unwrap()` returns `e.Err`. `NewOpError` returns `nil` when `err == nil`, else `&OpError{Op: op, Err: err}`.

---

## Hint 3

```go
func (e *OpError) Error() string {
    return e.Op + ": " + e.Err.Error()
}

func (e *OpError) Unwrap() error {
    return e.Err
}

func NewOpError(op string, err error) error {
    if err == nil {
        return nil
    }
    return &OpError{Op: op, Err: err}
}

func OpFromError(err error) string {
    var oe *OpError
    if errors.As(err, &oe) {
        return oe.Op
    }
    return ""
}

func MarkRetryable(err error, reason string) error {
    return &RetryableError{Reason: reason} // and set Unwrap to ErrRetryable
}
```

`RetryableError.Unwrap()` should return `ErrRetryable`. `RetryableError.Error()` can return `"retryable: " + reason`.
