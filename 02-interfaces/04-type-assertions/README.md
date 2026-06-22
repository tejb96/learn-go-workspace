# Type Assertions

## What you'll learn

- Type assertions and the comma-ok idiom
- Type switches on `any`
- Why `any` / `interface{}` should be a last resort

## Concept

### Plain English

An interface value holds a concrete type and value inside. A **type assertion** extracts the concrete type:

```go
v.(string)       // panics if v is not a string
s, ok := v.(string)  // ok is false if not a string — no panic
```

A **type switch** is the clean way to handle multiple possible types.

`any` (alias for `interface{}`) means "could be anything." Use it when the type truly varies at runtime (JSON decoding, fmt.Print, some generic containers). In application code, prefer **concrete types**, **interfaces you define**, or **generics** (Go 1.18+) so the compiler helps you.

### Go syntax

```go
func ParseAny(v any) (string, error) {
    switch x := v.(type) {
    case string:
        return x, nil
    case int:
        return strconv.Itoa(x), nil
    default:
        return "", ErrUnsupportedType
    }
}
```

## Annotated examples

```go
// WHY comma-ok: bare v.(int) crashes your server on bad input.
n, ok := v.(int)
if !ok {
    return 0, false
}
```

```go
// WHY any is last resort: you lose compile-time checking.
// Prefer func ParseInt(s string) (int, error) when input is always string.
func ParseAny(v any) (string, error) { ... }
```

## Common mistakes

- **Bare type assertion:** `v.(T)` panics — always use comma-ok in library code.
- **Using `any` everywhere:** Makes bugs runtime-only; narrow your API.
- **Asserting interface to wrong concrete type:** `var r io.Reader = strings.NewReader("x"); r.(*bytes.Buffer)` fails.
- **Confusing type assertion with type conversion:** `int64(5)` converts; `v.(int64)` asserts dynamic type.

## Further reading

- [Go spec — Type assertions](https://go.dev/ref/spec#Type_assertions)
- [Go blog — Generics](https://go.dev/blog/intro-generics) (alternative to `any`)

## API spec

| Symbol | Behavior |
|--------|----------|
| `ParseAny(v any)` | string/int/bool → string; else `ErrUnsupportedType` |
| `AsInt(v any)` | `(int, true)` for int/int64; else `(0, false)` |

## Before moving on

- [ ] I use comma-ok, not bare assertions
- [ ] I can write a type switch
- [ ] I know when **not** to use `any`
- [ ] All tests pass: `go test -v`
