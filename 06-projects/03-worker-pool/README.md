# Worker Pool Project

## What you'll learn

- Concurrent directory processing
- Worker pool pattern from Module 4 in a real task
- Collecting and sorting results

## Project brief

Implement `ProcessDir(dir string, workers int) ([]FileResult, error)`:

- List **regular files only** in `dir` (not subdirectories)
- Count lines per file (empty file → 0 lines; `"a\n"` → 1 line)
- Process with **worker pool** of size `workers`
- Return results sorted by **filename** (`Path` field = base name only)
- Error if directory missing or `workers < 1`

```go
type FileResult struct {
    Path  string
    Lines int
    Error string // optional, for unreadable files
}
```

## Common mistakes

- **Processing subdirectories:** Only top-level files.
- **Full path in Path:** Tests expect `"a.txt"`, not `/tmp/.../a.txt`.
- **Unsorted results:** Sort before return.
- **Deadlock:** Close `jobs` after enqueue; close `results` after workers finish.

## Further reading

- [Go blog — Pipelines and cancellation](https://go.dev/blog/pipelines)

## Before moving on

- [ ] `go test -v` passes
- [ ] `go test -race -v` passes
- [ ] I can explain worker count vs file count
