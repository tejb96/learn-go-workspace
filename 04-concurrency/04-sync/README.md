# sync Package

## What you'll learn

- `sync.WaitGroup` for goroutine completion
- `sync.Once` for one-time initialization
- Semaphore pattern with buffered channels to limit concurrency

## Concept

### Plain English

The **`sync`** package provides primitives when channels are awkward: mutexes, WaitGroups, Once, and more.

**WaitGroup** counts goroutines; call `Wait()` to block until the count hits zero.

**Once** guarantees a function runs exactly once — lazy initialization, singleton setup.

**Worker limits:** a buffered channel of empty structs acts as a semaphore — at most N goroutines hold a token at once.

### Go syntax

```go
var once sync.Once
var result int
once.Do(func() { result = expensive() })
```

```go
sem := make(chan struct{}, workers)
sem <- struct{}{} // acquire
// work
<-sem // release
```

## Annotated examples

```go
// WHY preserve order with out[i]: fan-out parallelism but deterministic output slice.
out[i] = fn(in[i])
```

## Common mistakes

- **WaitGroup Add after Wait started:** Add before launching goroutines.
- **Once inside hot path with side effects you expect repeated:** It runs once total.
- **Unbounded goroutines:** Always cap when work items >> CPU cores.
- **Closure loop variable bug:** Pass `i, v` as params to goroutine (Go 1.22+ loop var fix helps but explicit is clearer).

## Further reading

- [Package sync](https://pkg.go.dev/sync)

## API spec

| Function | Behavior |
|----------|----------|
| `ParallelMap(in, fn, workers)` | Ordered results; max `workers` concurrent fn calls |
| `OnceValue(fn)` | Returned func calls fn once, caches int result |

## Before moving on

- [ ] I can limit concurrency with a semaphore channel
- [ ] I use WaitGroup to wait for a batch of goroutines
- [ ] All tests pass: `go test -v`
