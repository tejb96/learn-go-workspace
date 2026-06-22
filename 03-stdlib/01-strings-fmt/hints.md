Try for at least 30 minutes before reading this.

## Hint 1

`strings` handles trimming, splitting, replacing, and case conversion. `fmt` helps with formatted output, but building lines with `strings.Join` is often clearer than repeated `fmt.Sprintf`.

---

## Hint 2

`Slugify`: `strings.TrimSpace`, `strings.ToLower`, replace spaces and `_` with `-`. Loop or use `strings.ReplaceAll` repeatedly to collapse `--` to `-`. `FormatTable`: first line joins headers with `" | "`, each row the same, lines joined with `\n` plus final `\n`.

---

## Hint 3

```go
func FormatTable(headers []string, rows [][]string) string {
    var b strings.Builder
    b.WriteString(strings.Join(headers, " | "))
    b.WriteByte('\n')
    for _, row := range rows {
        b.WriteString(strings.Join(row, " | "))
        b.WriteByte('\n')
    }
    return b.String()
}

func JoinLines(lines []string) string {
    return strings.Join(lines, "\n")
}
```

For slug collapsing, loop `strings.Contains(s, "--")` and replace until stable.
