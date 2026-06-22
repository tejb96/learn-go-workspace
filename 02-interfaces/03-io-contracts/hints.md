Try for at least 30 minutes before reading this.

## Hint 1

`io.Reader` and `io.Writer` are the core abstractions for streaming data. The same function can copy from a file, HTTP body, or `strings.NewReader` because all implement `Read`. Check the stdlib before reinventing — `io.CopyN` already exists.

---

## Hint 2

Your `CopyN` can delegate to `io.CopyN(dst, src, n)`. `ReadAll` should use `io.ReadAll(r)`. `WriteString` can use `io.WriteString(w, s)` or `fmt.Fprint`.

---

## Hint 3

```go
import "io"

func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error) {
    return io.CopyN(dst, src, n)
}

func ReadAll(r io.Reader) ([]byte, error) {
    return io.ReadAll(r)
}

func WriteString(w io.Writer, s string) (int, error) {
    return io.WriteString(w, s)
}
```

The lesson is understanding **why** one function works everywhere — read the README examples before copying.
