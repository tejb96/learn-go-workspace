# Functions and Multiple Return Values

## What you'll learn

- Declaring functions with multiple return values
- Variadic functions (`...T`)
- Returning errors instead of panicking on invalid input
- Named return values (and why to use them sparingly)

## Concept

### Plain English

Go functions can return more than one value. The most common pattern is `(result, error)` — success data plus a signal for failure. Callers write `v, err := Foo()` and check `err` before using `v`.

Functions can also accept a variable number of arguments with **variadic** parameters: `func Sum(nums ...int)`. Inside the function, `nums` is a slice.

When something goes wrong in a function like division by zero, **return an error** rather than letting the program panic. Panics are for truly broken program state, not expected validation failures.

### Go syntax

```go
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divide by zero")
    }
    return a / b, nil
}

func Sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

**Named return values** (optional):

```go
func SplitHostPort(hostport string) (host string, port int, err error) {
    // host, port, err are pre-declared; bare return sets them
    return // rarely used except small functions; can hurt readability
}
```

Prefer explicit `return value, nil` in learning code until you know when naked returns help.

## Annotated examples

```go
// WHY multiple returns: forces the caller to handle failure at the call site.
// Compare to exceptions that can bubble up unnoticed.
func ReadConfig(path string) ([]byte, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err // nil slice is the zero value — safe to return on error
    }
    return data, nil
}
```

```go
// WHY variadic: fmt.Println accepts any number of args the same way.
func LogParts(parts ...string) {
    for i, p := range parts {
        if i > 0 {
            fmt.Print(" | ")
        }
        fmt.Print(p)
    }
}
```

## Common mistakes

- **Panicking on bad input:** `panic("divide by zero")` makes tests and callers harder to reason about. Return `error`.
- **Using named returns everywhere:** Bare `return` in long functions hides what is being returned.
- **Forgetting variadic args become a slice:** `Sum(1, 2, 3)` and `Sum([]int{1,2,3}...)` are equivalent.
- **Returning wrong zero value on error:** On failure, return the type's zero value (`0`, `nil`, `""`) plus the error.

## Further reading

- [Go spec — Function declarations](https://go.dev/ref/spec#Function_declarations)
- [Effective Go — Errors](https://go.dev/doc/effective_go#errors)

## API spec

Implement in `solution.go`:

| Function | Signature | Behavior |
|----------|-----------|----------|
| `Sum` | `(nums ...int) int` | Sum all arguments; empty → `0` |
| `Divide` | `(a, b int) (int, error)` | Integer division; error if `b == 0` |

## Before moving on

- [ ] I can write a variadic function and call it with zero or many args
- [ ] I return `(zeroValue, err)` on failure instead of panicking
- [ ] I understand why `(T, error)` is idiomatic
- [ ] All tests pass: `go test -v`
