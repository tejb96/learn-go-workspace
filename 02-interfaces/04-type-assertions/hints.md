Try for at least 30 minutes before reading this.

## Hint 1

`any` is an alias for `interface{}` — it holds any dynamic type. Use a **type switch** to branch on concrete types. The comma-ok form `v.(int)` returns whether the assertion succeeded without panicking.

---

## Hint 2

```go
switch x := v.(type) {
case string:
    return x, nil
case int:
    return strconv.Itoa(x), nil
// ...
default:
    return "", ErrUnsupportedType
}
```

For `AsInt`, try `v.(int)` and `v.(int64)` with comma-ok.

---

## Hint 3

```go
func ParseAny(v any) (string, error) {
    switch x := v.(type) {
    case string:
        return x, nil
    case int:
        return strconv.Itoa(x), nil
    case bool:
        return strconv.FormatBool(x), nil
    default:
        return "", ErrUnsupportedType
    }
}

func AsInt(v any) (int, bool) {
    if n, ok := v.(int); ok {
        return n, true
    }
    if n, ok := v.(int64); ok {
        return int(n), true
    }
    return 0, false
}
```

Import `strconv`. Never use bare `v.(int)` in production without ok — it panics on mismatch.
