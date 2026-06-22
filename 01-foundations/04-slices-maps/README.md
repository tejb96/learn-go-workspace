# Slices and Maps

## What you'll learn

- Slice header: pointer, length, capacity
- `append` behavior and when reallocation happens
- `copy` and slicing without corrupting shared state
- Map basics and safe nil map reads

## Concept

### Plain English

A **slice** is a view into an underlying array. It has a **length** (how many elements you see) and **capacity** (how many slots exist in the backing array from the slice's start index). `append` adds elements; if capacity is full, Go allocates a larger array and copies data.

Because slices share backing arrays, `b := a` and subslices like `a[:2]` can surprise you: mutating one may mutate another unless you limit capacity with a **full slice expression** `a[low:high:max]`.

A **map** is a hash table. The zero value is `nil`; reading a missing key returns the zero value for the value type. Writing to a nil map panics — initialize with `make(map[K]V)`.

### Go syntax

```go
s := make([]int, 0, 10) // len 0, cap 10
s = append(s, 1, 2, 3)

sub := s[1:3]   // len 2, shares backing array with s
sub[0] = 99     // may change s[1]

// Limit capacity so append on sub does not overwrite s:
sub := s[1:3:3]

dst := make([]int, len(src))
copy(dst, src)
```

```go
m := make(map[string]int)
m["key"] = 42
v, ok := m["key"] // ok is false if missing
```

### Slice header (mental model)

```
  slice { ptr, len, cap }
            |
            v
  [ | | | | | ]  backing array
    ^     ^
    |     len elements visible
    cap slots available from ptr
```

## Annotated examples

```go
// WHY append may not change other slices: reallocation when cap exhausted.
a := []int{1, 2}
b := append(a, 3) // if cap was 2, new array; a unchanged
```

```go
// WHY full slice expression: prevents append from clobbering parent.
orig := []int{1, 2, 3, 4}
window := orig[1:3:3] // cap=2, append reallocates
window = append(window, 99)
// orig stays [1 2 3 4]
```

## Common mistakes

- **Assuming append always returns a new slice:** It may reuse the same backing array if capacity allows.
- **Using nil slice vs empty slice:** Both have len 0; JSON and some APIs treat them differently — `[]int{}` is non-nil empty.
- **Writing to nil map:** `var m map[string]int; m["x"] = 1` panics.
- **Forgetting `copy` length:** `copy(dst, src)` copies `min(len(dst), len(src))` elements.

## Further reading

- [Go spec — Slice types](https://go.dev/ref/spec#Slice_types)
- [Go blog — Arrays, slices (and strings and 'runes')](https://go.dev/blog/slices)

## API spec

| Function | Behavior |
|----------|----------|
| `AppendUnique` | Append `v` only if not already in `s` |
| `SafeSlice` | Slice with clamped bounds; invalid range → `[]int{}` |
| `MapInvert` | Swap keys/values; duplicate values → last key wins; nil input → empty map |
| `CapAfterAppend` | `cap` after `append(s, 0)` |

## Before moving on

- [ ] I can explain len vs cap
- [ ] I know when append reallocates
- [ ] I can invert a map safely
- [ ] All tests pass: `go test -v`
