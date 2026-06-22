Try for at least 30 minutes before reading this.

## Hint 1

`os.ReadDir` lists non-recursive entries. Skip directories. Send file paths on a `jobs` channel; workers count lines with `bufio.Scanner` and send `FileResult` on `results` channel.

---

## Hint 2

Clamp `workers` to at least 1 or error if 0. Collect results, `sort.Slice` by `Path`. Use `filepath.Base` for `Path` in results (filename only, not full path).

---

## Hint 3

```go
jobs := make(chan string)
results := make(chan FileResult)
var wg sync.WaitGroup
for w := 0; w < workers; w++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        for path := range jobs {
            results <- countLines(path)
        }
    }()
}
go func() {
    wg.Wait()
    close(results)
}()
// enqueue files, close(jobs), collect from results
```

Return error if `workers < 1` or dir missing.
