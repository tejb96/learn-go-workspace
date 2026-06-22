# Variables and Zero Values

## What you'll learn

- How Go initializes variables you do not explicitly set (zero values)
- When to use `var` versus short declaration `:=`
- Why Go functions commonly return multiple values

## Concept

### Plain English

Every variable in Go has a type, and every type has a **zero value** — the value Go uses when you declare a variable without assigning anything. You never get "uninitialized memory" like in C. An `int` starts at `0`, a `bool` at `false`, a `string` at `""` (empty string, not null), and slices, maps, channels, and pointers start at `nil`.

Go also lets functions return **multiple values** in one return statement. This is not an afterthought — it is how the language handles errors (`result, err := doThing()`) and how you swap or split data without out-parameters.

### Go syntax

**Declaration with zero value (`var`):**

```go
var count int     // count is 0
var name string   // name is ""
var ok bool       // ok is false
```

**Short declaration (`:=`) — only inside functions:**

```go
func example() {
    x := 10        // int, inferred
    y := "hello"   // string, inferred
    // var z int = 10  // equivalent to x, but := is idiomatic in functions
}
```

**Why `var` at package level or when you want an explicit type:**

```go
var DefaultPort = 8080 // package-level; type inferred as int

func parse() (string, error) {
    var buf strings.Builder // zero value is usable; Builder starts empty
    // ...
    return buf.String(), nil
}
```

**Multiple return values:**

```go
func Swap(a, b string) (string, string) {
    return b, a // idiomatic swap without a temp variable
}
```

## Annotated examples

```go
// WHY check empty string this way: "" is the zero value for string.
// There is no nil string in Go.
func Greet(name string) string {
    if name == "" {
        return "Hello, stranger!"
    }
    return "Hello, " + name + "!"
}

// WHY return (value, error): callers must acknowledge failure explicitly.
// Ignoring err with _ is a conscious choice, not an accident.
func ParsePort(s string) (int, error) {
    n, err := strconv.Atoi(s)
    if err != nil {
        return 0, err // return zero value for int on failure
    }
    return n, nil
}
```

```go
// WHY var at package scope: := cannot be used outside function bodies.
var Version = "1.0.0"

// WHY := inside functions: less noise when type is obvious from the RHS.
func double(n int) int {
    result := n * 2
    return result
}
```

## Common mistakes

- **Assuming `""` means "missing" like `null`:** Empty string is a valid value. Use pointers or `sql.NullString` when you need three states (unset / empty / set).
- **Using `:=` when you need to assign to an existing variable:** `:=` declares at least one new variable. Use `=` for reassignment.
- **Shadowing with `:=`:** `x := 1` inside an inner block can hide an outer `x`. The compiler allows it; your brain may not.
- **Ignoring the second return value:** `value, _ := risky()` hides errors. Only do this when failure is impossible or irrelevant.

## Further reading

- [Go spec — Zero values](https://go.dev/ref/spec#The_zero_value)
- [Effective Go — Declarations](https://go.dev/doc/effective_go#declarations)

## API spec

Implement in `solution.go`:

| Function | Signature | Behavior |
|----------|-----------|----------|
| `Swap` | `(a, b string) (string, string)` | Return `(b, a)` |
| `DescribeType` | `(kind string) (string, error)` | For `"string"`, `"int"`, `"bool"`, `"slice"` return the exact zero-value description strings tested; otherwise error |
| `DefaultGreeting` | `(name string) string` | Empty name → `"Hello, stranger!"`; else `"Hello, <name>!"` |

## Before moving on

- [ ] I can name the zero value for `int`, `string`, `bool`, and `slice`
- [ ] I know when `var` is required vs when `:=` is idiomatic
- [ ] I can write a function that returns two values
- [ ] All tests pass: `go test -v`
