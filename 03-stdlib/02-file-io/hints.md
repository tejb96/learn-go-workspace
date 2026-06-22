Try for at least 30 minutes before reading this.

## Hint 1

`os.ReadFile` reads an entire file. Split on `\n` and handle the last empty element from a trailing newline. Atomic writes mean: write to a temp file in the same directory, `Sync`, then `Rename` over the target.

---

## Hint 2

`ReadLines`: `bytes.Split` or `strings.Split` after `os.ReadFile`, trim trailing empty line from final `\n`. Return `[]string{}` for empty file. `WriteAtomic`: `os.MkdirAll` on parent, `os.CreateTemp` in same dir, write, `Close`, `os.Rename`.

---

## Hint 3

```go
func ReadLines(path string) ([]string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    if len(data) == 0 {
        return []string{}, nil
    }
    text := string(data)
    text = strings.TrimSuffix(text, "\n")
    if text == "" {
        return []string{}, nil
    }
    return strings.Split(text, "\n"), nil
}
```

For `FileExists`, use `os.Stat` and check `err == nil`.
