# Interface Basics

## What you'll learn

- What an interface is in Go
- Implicit satisfaction (no `implements` keyword)
- How interfaces enable polymorphism
- How this differs from Java and C#

## Concept

### Plain English

An **interface** is a contract: a set of method names and signatures. Any type that has those methods **automatically** satisfies the interface. You do not declare conformance — the compiler checks it at compile time when you pass a value to a function expecting that interface.

This is **implicit satisfaction**. In Java or C#, you write `class Dog implements Speaker`. In Go, if `Dog` has `Speak() string`, it already **is** a `Speaker` wherever one is needed.

Interfaces let you write functions that work on **behavior**, not concrete types: `Greet(s Speaker)` accepts a `Person`, `Robot`, or any future type with `Speak()`.

### Go syntax

```go
type Speaker interface {
    Speak() string
}

func Greet(s Speaker) string {
    return "Hello, " + s.Speak()
}
```

The interface value holds a **dynamic type** and **dynamic value** (a pair). A `nil` interface is different from an interface holding a nil pointer — a subtle gotcha for later.

### Go vs Java / C#

| | Go | Java / C# |
|---|-----|-----------|
| Declare conformance | Never — implicit | `implements` / `: IInterface` |
| Add interface to existing type | Yes, if methods match | Often requires editing the class |
| Interface size | Usually small (1–3 methods) | Often large |

Go proverb: **"The bigger the interface, the weaker the abstraction."**

## Annotated examples

```go
// WHY accept Speaker not Person: Greet works on any future Speaker
// without changing this function.
func Greet(s Speaker) string {
    return "Greeting: " + s.Speak()
}

// WHY value receiver on Person: Speak reads fields, does not mutate.
func (p Person) Speak() string {
    return "Hi, I'm " + p.Name
}
```

```go
// WHY slice of interface: heterogeneous collection — Person and Robot together.
speakers := []Speaker{
    Person{Name: "Ada"},
    Robot{Model: "X1"},
}
```

## Common mistakes

- **Empty interface overuse:** Prefer small, focused interfaces like `Speaker`.
- **Assuming nil pointer satisfies interface meaningfully:** `var s Speaker = (*Person)(nil)` is not nil as an interface value.
- **Pointer vs value receiver:** Only the receiver type you define methods on satisfies the interface (`Person` vs `*Person`).
- **Checking concrete type when interface suffices:** Accept `Speaker`, not `Person`, when behavior is all you need.

## Further reading

- [Go spec — Interface types](https://go.dev/ref/spec#Interface_types)
- [Effective Go — Interfaces](https://go.dev/doc/effective_go#interfaces)

## API spec

| Symbol | Behavior |
|--------|----------|
| `Speaker` | interface with `Speak() string` |
| `Person.Speak()` | `"Hi, I'm " + Name` |
| `Robot.Speak()` | `"Beep boop, model " + Model` |
| `Greet(s Speaker)` | `"Greeting: " + s.Speak()` |
| `SpeakAll([]Speaker)` | Each `Speak()` in order; nil input → empty slice |

## Before moving on

- [ ] I can explain implicit interface satisfaction
- [ ] I can contrast Go with Java/C# interface declaration
- [ ] I wrote a function that accepts an interface type
- [ ] All tests pass: `go test -v`
