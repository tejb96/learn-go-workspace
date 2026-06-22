Try for at least 30 minutes before reading this.

## Hint 1

Read `helper_test.go` — `t.Helper()` marks a function so failures report the caller's line, not the helper's. Tests use `t.Run` for subtests. Implement logic in `solution.go` only.

---

## Hint 2

Palindrome: compare s with reversed runes or two-pointer from both ends. `SumPositive`: loop, add if `n > 0`.

---

## Hint 3

```go
func IsPalindrome(s string) bool {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}

func SumPositive(nums []int) int {
    sum := 0
    for _, n := range nums {
        if n > 0 {
            sum += n
        }
    }
    return sum
}
```

Run subtests: `go test -v -run TestIsPalindrome/classic`
