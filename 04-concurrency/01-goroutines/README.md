# Goroutines

## What you'll learn

- Starting goroutines with `go`
- Why shared memory without locks is unsafe
- Using `sync.Mutex` and `sync.WaitGroup`
- Finding races with `go test -race`

## Concept

### Plain English

A **goroutine** is a lightweight thread managed by the Go runtime. Start one with `go f()`. Thousands are normal.

When two goroutines update the same variable without coordination, you have a **data race** — undefined behavior. The result is nondeterministic: tests may flake, counts may be wrong, production may corrupt data.

The **race detector** (`go test -race`) instruments your code and reports races at runtime. Use it on any package with concurrency.

### Go syntax

```go
go func() {
    doWork()
}()

var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    work()
}()
wg.Wait()
```

```go
var mu sync.Mutex
mu.Lock()
counter++
mu.Unlock()
```

## Data race example

`UnsafeCounter` in `solution.go` increments without a lock. Run:

```bash
go test -v
go test -race -v
```

`-race` should report a race on `UnsafeCounter` when tests run concurrent increments. `SafeCounter` with a mutex should pass both.

## Annotated examples

```go
// WHY mutex around n++: read-modify-write is three steps; another goroutine can interleave.
func (c *SafeCounter) Inc() {
    c.mu.Lock()
    c.n++
    c.mu.Unlock()
}
```

```go
// WHY WaitGroup: main goroutine must wait for workers to finish before checking results.
wg.Add(1)
go func() {
    defer wg.Done()
    // ...
}()
wg.Wait()
```

## Common mistakes

- **Forgetting `wg.Add` before `go`:** Race with `Wait`.
- **Copying WaitGroup:** Pass pointer or keep in one scope.
- **Locking after unlock path panics:** Use `defer mu.Unlock()` after Lock.
- **Assuming `go` runs immediately:** Scheduling is nondeterministic.

## Further reading

- [Go blog — Share by communicating](https://go.dev/blog/codelab-share)
- [Race detector](https://go.dev/doc/articles/race_detector)

## API spec

| Symbol | Behavior |
|--------|----------|
| `UnsafeCounter` | Provided — intentionally racy |
| `SafeCounter` | Mutex-protected `Inc` / `Value` |
| `RunAll(fns)` | Run each func in its own goroutine; wait for all |

## Before moving on

- [ ] I ran `go test -race` and understand the UnsafeCounter report
- [ ] `SafeCounter` passes concurrent test with exact count
- [ ] `RunAll` waits for all goroutines
- [ ] All tests pass: `go test -v`
