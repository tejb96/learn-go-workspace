# strings and fmt

## What you'll learn

- Common `strings` operations: trim, join, replace, case
- Building text with `strings.Builder` vs `fmt`
- When formatting is data transformation vs presentation

## Concept

### Plain English

The **`strings`** package is your toolkit for text manipulation without regex. The **`fmt`** package formats values for humans — printing, sprintf, errors. For slug generation and table building, `strings` is usually the right layer; use `fmt` when you need verbs like `%d`, `%q`, or `%v`.

### Go syntax

```go
s := strings.TrimSpace("  hello  ")
parts := strings.Fields("a  b")          // split on whitespace
joined := strings.Join(parts, "-")
lower := strings.ToLower("Go")
replaced := strings.ReplaceAll(s, " ", "-")
```

```go
var b strings.Builder
b.WriteString("name")
b.WriteByte('\n')
result := b.String()
```

## Annotated examples

```go
// WHY ToLower before replace: slugs are conventionally lowercase URLs.
func Slugify(s string) string {
    s = strings.TrimSpace(strings.ToLower(s))
    // ...
}
```

```go
// WHY Builder over += in loops: avoids O(n²) string copying.
func FormatTable(headers []string, rows [][]string) string {
    var b strings.Builder
    // ...
    return b.String()
}
```

## Common mistakes

- **Using `fmt` for simple joins:** `strings.Join` is clearer than `fmt.Sprintf("%s,%s", a, b)` for many cases.
- **Mutating strings:** Strings are immutable; every operation returns a new string.
- **Forgetting Unicode:** `ToLower` handles most Unicode; do not hand-roll ASCII-only unless required.
- **Trailing newlines:** Be explicit about whether output ends with `\n` — tests often care.

## Further reading

- [Package strings](https://pkg.go.dev/strings)
- [Package fmt](https://pkg.go.dev/fmt)

## API spec

| Function | Behavior |
|----------|----------|
| `Slugify(s)` | Trim, lower, spaces/`_` → `-`, collapse `--` |
| `FormatTable(headers, rows)` | Pipe-separated columns, newline after each row |
| `JoinLines(lines)` | Join with `\n`; empty input → `""` |

## Before moving on

- [ ] I reach for `strings` before regex for simple transforms
- [ ] I know when `strings.Builder` beats concatenation
- [ ] All tests pass: `go test -v`
