Try for at least 30 minutes before reading this.

## Hint 1

Variadic parameters (`nums ...int`) arrive as a slice inside the function. An empty slice has length 0, and summing nothing should give you the zero value for `int`. For division by zero, Go does not throw — you return an error as the second value.

---

## Hint 2

Loop over `nums` with `for _, n := range nums` and accumulate. Define a package-level sentinel like `var ErrDivideByZero = errors.New("divide by zero")` and return `(0, ErrDivideByZero)` when `b == 0`. Integer division in Go truncates toward zero automatically.

---

## Hint 3

```go
var ErrDivideByZero = errors.New("divide by zero")

func Sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, ErrDivideByZero
    }
    return a / b, nil
}
```

The divide-by-zero test only checks that an error is returned, not a specific sentinel — but using one is good practice for later lessons.
