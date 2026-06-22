Try for at least 30 minutes before reading this.

## Hint 1

Interface embedding in a type definition means the outer interface includes all methods of the inner ones. `ReadWriter` is satisfied only when a type has **both** `Read` and `Write`. Use type assertions with the comma-ok form to check which interfaces a value satisfies.

---

## Hint 2

In `Describe`, check `ReadWriter` first, then `Reader`, then `Writer`:

```go
if _, ok := v.(ReadWriter); ok { return "readwriter" }
```

Order matters — a ReadWriter is also a Reader, so test the combined interface first.

---

## Hint 3

```go
func Describe(v any) string {
    if _, ok := v.(ReadWriter); ok {
        return "readwriter"
    }
    if _, ok := v.(Reader); ok {
        return "reader"
    }
    if _, ok := v.(Writer); ok {
        return "writer"
    }
    return "none"
}

func Duplicate(r Reader, writers ...Writer) []int {
    data := r.Read()
    counts := make([]int, len(writers))
    for i, w := range writers {
        counts[i] = w.Write(data)
    }
    return counts
}
```
