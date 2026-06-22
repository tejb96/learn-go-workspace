# File I/O

## What you'll learn

- Reading and writing files with `os`
- Line-oriented file processing
- Atomic writes with temp file + rename

## Concept

### Plain English

Files on disk are bytes. `os.ReadFile` and `os.WriteFile` are the simple whole-file helpers. For production writes, **atomic replace** avoids half-written files if the process crashes mid-write: write to a temporary file in the same directory, fsync, then `rename` over the destination (rename is atomic on POSIX).

Always handle errors from `os` operations — missing files, permission denied, and disk full are normal failure modes.

### Go syntax

```go
data, err := os.ReadFile("config.yaml")
if err != nil {
    return err
}

err = os.WriteFile("out.txt", data, 0o644)
```

```go
dir := filepath.Dir(path)
f, err := os.CreateTemp(dir, ".tmp-*")
// write, f.Close(), os.Rename(f.Name(), path)
```

## Annotated examples

```go
// WHY rename in same directory: POSIX guarantees atomic replace.
func WriteAtomic(path string, data []byte) error {
    // temp in filepath.Dir(path) → Rename → path
}
```

```go
// WHY TrimSuffix on final newline: "a\nb\n" → ["a","b"], not ["a","b",""].
func ReadLines(path string) ([]string, error) { ... }
```

## Common mistakes

- **Ignoring file permissions:** `0o644` vs `0o600` matters for secrets.
- **Not closing files:** Prefer `defer f.Close()` when streaming; `WriteFile` handles it.
- **Writing directly to final path:** Crash leaves corrupted file — use atomic pattern.
- **Using `\` on Windows paths:** Prefer `filepath.Join`.

## Further reading

- [Package os](https://pkg.go.dev/os)
- [Package path/filepath](https://pkg.go.dev/path/filepath)

## API spec

| Function | Behavior |
|----------|----------|
| `ReadLines(path)` | Lines without `\n`; empty file → `[]string{}`; error if missing |
| `WriteAtomic(path, data)` | Create dirs, atomic write |
| `FileExists(path)` | `true` if `os.Stat` succeeds |

## Before moving on

- [ ] I can read a whole file and split lines
- [ ] I understand temp + rename for atomic writes
- [ ] All tests pass: `go test -v`
