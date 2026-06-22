Try for at least 30 minutes before reading this.

## Hint 1

A slice is a small header (pointer, len, cap) over an array. `append` may write into spare capacity or allocate a new array. Two slices can share a backing array if created with `s[:n:n]` limiting cap matters.

---

## Hint 2

`AppendUnique`: loop to check duplicates, then `append(s, v)`. `SafeSlice`: clamp indices, use `s[low:high]` or return `[]int{}`. `MapInvert`: `make(map[int]string)` and range over `m`. `CapAfterAppend`: `s2 := append(s, 0); return cap(s2)`.

---

## Hint 3

```go
func SafeSlice(s []int, low, high int) []int {
    if low < 0 {
        low = 0
    }
    if high > len(s) {
        high = len(s)
    }
    if low > high {
        return []int{}
    }
    return s[low:high]
}

func CapAfterAppend(s []int) int {
    s2 := append(s, 0)
    return cap(s2)
}
```

For `MapInvert` on nil input, return `map[int]string{}` not nil if tests require empty map — check DeepEqual with empty map literal.
