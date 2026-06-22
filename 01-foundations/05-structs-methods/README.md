# Structs and Methods

## What you'll learn

- Defining structs and constructing values
- Methods with value vs pointer receivers
- Constructor-style functions returning `(*T, error)`

## Concept

### Plain English

A **struct** groups named fields into one type. Methods attach behavior to a type via a **receiver** — like a function belonging to the struct.

Use a **value receiver** when the method only reads the struct or needs a copy. Use a **pointer receiver** when the method mutates the struct, when the struct is large, or when consistency matters (if one method uses `*T`, often all methods on that type do too).

Constructor helpers like `NewRectangle` are idiomatic when validation or opaque initialization is needed. They return a pointer and an error.

### Go syntax

```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r *Rectangle) Scale(f float64) {
    r.Width *= f
    r.Height *= f
}
```

```go
r := Rectangle{Width: 3, Height: 4}
r.Scale(2) // Go passes &r automatically for pointer receiver

p := &Rectangle{Width: 1, Height: 1}
p.Scale(2)
```

## Annotated examples

```go
// WHY pointer receiver on Scale: value receiver would copy Rectangle;
// mutations would die when the method returns.
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

```go
// WHY constructor returns error: invalid dimensions are expected input failures.
func NewRectangle(w, h float64) (*Rectangle, error) {
    if w < 0 || h < 0 {
        return nil, errors.New("dimensions must be non-negative")
    }
    return &Rectangle{Width: w, Height: h}, nil
}
```

## Common mistakes

- **Pointer receiver but calling on unaddressable value:** Rare edge cases with maps/slices of structs.
- **Mixing value and pointer receivers inconsistently:** Confusing for maintainers; pick a convention per type.
- **Forgetting `New` can return `(nil, err)`:** Callers must check `err` before using the pointer.
- **Comparing structs with `==`:** Only if all fields are comparable; slices/maps inside struct make it illegal.

## Further reading

- [Go spec — Struct types](https://go.dev/ref/spec#Struct_types)
- [Effective Go — Methods](https://go.dev/doc/effective_go#methods)

## API spec

| Symbol | Behavior |
|--------|----------|
| `Rectangle` | `Width`, `Height` float64 |
| `(r Rectangle) Area()` | `Width * Height` |
| `(r *Rectangle) Scale(factor)` | Multiply both dimensions in place |
| `NewRectangle(w, h)` | Error if either negative; else `&Rectangle{w,h}` |

## Before moving on

- [ ] I know when to use value vs pointer receivers
- [ ] I understand why `Scale` must use `*Rectangle`
- [ ] I can write a small constructor with validation
- [ ] All tests pass: `go test -v`

## Next module

When all foundation lessons pass, continue to **02-interfaces/01-basics**.
