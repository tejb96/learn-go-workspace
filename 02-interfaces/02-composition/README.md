# Interface Composition

## What you'll learn

- Embedding interfaces inside other interfaces
- Building larger contracts from smaller ones
- Checking which interfaces a value satisfies

## Concept

### Plain English

Go lets you **embed** interfaces inside other interfaces. If `ReadWriter` embeds `Reader` and `Writer`, any type that satisfies `ReadWriter` must implement **all** methods from both.

This mirrors struct embedding: compose small pieces into larger abstractions instead of one giant interface.

The standard library uses this heavily — `io.ReadWriter` embeds `io.Reader` and `io.Writer`.

### Go syntax

```go
type Reader interface {
    Read() []byte
}

type Writer interface {
    Write(p []byte) int
}

type ReadWriter interface {
    Reader
    Writer
}
```

A concrete type satisfies `ReadWriter` only when it has both methods. It may satisfy `Reader` alone without `Writer`.

## Annotated examples

```go
// WHY embed: ReadWriter is exactly "can read AND write" — no duplicate method lists.
type ReadWriter interface {
    Reader
    Writer
}
```

```go
// WHY check ReadWriter before Reader: every ReadWriter is a Reader;
// labeling it "reader" alone would under-report capability.
func Describe(v any) string {
    if _, ok := v.(ReadWriter); ok {
        return "readwriter"
    }
    // ...
}
```

## Common mistakes

- **Checking Reader before ReadWriter:** Combined interface must be tested first.
- **Huge embedded stacks:** Prefer stdlib patterns — small composable interfaces.
- **Embedding in structs vs interfaces:** Struct embedding promotes methods; interface embedding unions method sets.

## Further reading

- [Go spec — Interface types (embedded methods)](https://go.dev/ref/spec#Interface_types)
- [Package io — ReadWriter](https://pkg.go.dev/io#ReadWriter)

## API spec

| Symbol | Behavior |
|--------|----------|
| `ReadWriter` | embeds `Reader` and `Writer` |
| `Describe(v any)` | `"readwriter"`, `"reader"`, `"writer"`, or `"none"` |
| `Duplicate(r, writers...)` | Read once, write same bytes to each writer; return byte counts |

## Before moving on

- [ ] I can embed interfaces to compose larger ones
- [ ] I know to check broader interfaces before narrower ones
- [ ] All tests pass: `go test -v`
