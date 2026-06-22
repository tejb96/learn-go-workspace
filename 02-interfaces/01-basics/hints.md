Try for at least 30 minutes before reading this.

## Hint 1

In Go, a type satisfies an interface when it has the required methods — you never write `implements Speaker`. Your `Person` and `Robot` types need `Speak() string` methods with the exact strings the tests expect.

---

## Hint 2

`Person.Speak` returns `"Hi, I'm " + p.Name`. `Robot.Speak` returns `"Beep boop, model " + r.Model`. `Greet` takes a `Speaker` and returns `"Greeting: " + s.Speak()`. For `SpeakAll`, preallocate or append in a loop.

---

## Hint 3

```go
func (p Person) Speak() string {
    return "Hi, I'm " + p.Name
}

func Greet(s Speaker) string {
    return "Greeting: " + s.Speak()
}

func SpeakAll(speakers []Speaker) []string {
    out := make([]string, 0, len(speakers))
    for _, s := range speakers {
        out = append(out, s.Speak())
    }
    return out
}
```

For `SpeakAll(nil)`, ranging over nil is fine — return `out` which is an empty non-nil slice from `make`.
