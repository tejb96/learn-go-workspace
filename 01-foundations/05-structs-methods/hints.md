Try for at least 30 minutes before reading this.

## Hint 1

Methods are functions with a receiver: `(r Rectangle)` or `(r *Rectangle)`. Value receivers get a copy; pointer receivers can modify the original. Use pointers when the method mutates state or when the struct is large.

---

## Hint 2

`Area` does not need to mutate — value receiver on `Rectangle` is fine. `Scale` must use `*Rectangle` and multiply fields in place. `NewRectangle` validates inputs and returns `(&Rectangle{...}, nil)` or `(nil, errors.New(...))`.

---

## Hint 3

```go
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func NewRectangle(width, height float64) (*Rectangle, error) {
    if width < 0 || height < 0 {
        return nil, errors.New("dimensions must be non-negative")
    }
    return &Rectangle{Width: width, Height: height}, nil
}
```

Go automatically takes address for `r.Scale` when `r` is a value — but mutation only works with pointer receiver.
