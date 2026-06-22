# io.Reader and io.Writer

## What you'll learn

- The `io.Reader` and `io.Writer` interfaces
- Why one function can work on files, HTTP bodies, and buffers
- When to delegate to the standard library

## Concept

### Plain English

Most I/O in Go is streaming: read or write chunks of bytes without loading everything into memory at once. Two tiny interfaces power this:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Dozens of types implement these — files, network connections, HTTP response bodies, `bytes.Buffer`, `strings.Reader`, compressors, hashers. A function accepting `io.Reader` does not care **where** bytes come from.

That is the power of small interfaces: **one abstraction, many sources and sinks**.

### Same function, three sources

```go
func CopyN(dst io.Writer, src io.Reader, n int64) (int64, error) {
    return io.CopyN(dst, src, n)
}

// File on disk
f, _ := os.Open("data.txt")
CopyN(&buf, f, 1024)

// HTTP response body (client lesson builds on this)
resp, _ := http.Get(url)
CopyN(&buf, resp.Body, 1024)

// In-memory string
CopyN(&buf, strings.NewReader("hello"), 5)
```

## Annotated examples

```go
// WHY io.Reader not *os.File: tests can pass strings.NewReader;
// production can pass resp.Body — same code path.
func ReadAll(r io.Reader) ([]byte, error) {
    return io.ReadAll(r)
}
```

```go
// WHY bytes.Buffer for dst: it implements io.Writer — collect output in memory in tests.
dst := &bytes.Buffer{}
io.CopyN(dst, src, n)
```

## Common mistakes

- **Reading entire files with `ioutil` deprecated patterns:** Use `os.ReadFile` for whole files, `io.Reader` for streams.
- **Forgetting to close `resp.Body`:** Not this lesson, but HTTP bodies are Readers that must close.
- **Ignoring partial Read:** `Read` may return `n > 0` and `err == io.EOF` on the last chunk.
- **Reimplementing io.CopyN:** Delegate to stdlib unless you are learning how it works.

## Further reading

- [Package io](https://pkg.go.dev/io)
- [Effective Go — Interfaces and io](https://go.dev/doc/effective_go#interfaces_and_types)

## API spec

| Function | Behavior |
|----------|----------|
| `CopyN(dst, src, n)` | Copy at most n bytes (stdlib semantics) |
| `ReadAll(r)` | Read until EOF |
| `WriteString(w, s)` | Write string to w |

## Before moving on

- [ ] I can name three types that implement `io.Reader`
- [ ] I understand why `CopyN` accepts interfaces, not concrete file types
- [ ] All tests pass: `go test -v`
