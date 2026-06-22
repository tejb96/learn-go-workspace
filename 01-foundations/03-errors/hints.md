Try for at least 30 minutes before reading this.

## Hint 1

Sentinel errors are package-level variables compared with `errors.Is`. Wrapping adds context but preserves the chain when you use `fmt.Errorf` with `%w`. Custom error types implement the `error` interface and can be inspected with `errors.As`.

---

## Hint 2

`Wrap` should be `return fmt.Errorf("%s: %w", msg, err)`. In `Find`, check `id == ""` first and return `&ValidationError{Field: "id", Code: "required"}`. Then look up the map; if the key is missing, return `""` and `ErrNotFound`.

---

## Hint 3

```go
func Wrap(err error, msg string) error {
    if err == nil {
        return nil
    }
    return fmt.Errorf("%s: %w", msg, err)
}

func Find(store map[string]string, id string) (string, error) {
    if id == "" {
        return "", &ValidationError{Field: "id", Code: "required"}
    }
    v, ok := store[id]
    if !ok {
        return "", ErrNotFound
    }
    return v, nil
}
```

Import `fmt` in solution.go. The wrapped error message must be exactly `lookup user: not found` when wrapping `ErrNotFound` with message `lookup user`.
