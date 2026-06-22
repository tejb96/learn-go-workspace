Try for at least 30 minutes before reading this.

## Hint 1

Go gives every variable a **zero value** when you declare it without an initializer. For `string` that is `""`, not `null` like some other languages. Multiple return values are normal — look at how other stdlib functions return `(value, error)`.

---

## Hint 2

For `Swap`, you do not need a temp variable, but you can use one. Return `(b, a)` directly. For `DescribeType`, use a `switch` on `kind` and return `errors.New` for unknown kinds. Compare empty string with `name == ""`.

---

## Hint 3

```go
func Swap(a, b string) (string, string) {
    return b, a
}

func DefaultGreeting(name string) string {
    if name == "" {
        return "Hello, stranger!"
    }
    return "Hello, " + name + "!"
}
```

For `DescribeType`, map each kind to the exact strings the tests expect, including backticks around the empty string in the string case.
