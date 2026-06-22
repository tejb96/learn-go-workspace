# Concurrency Patterns

## What you'll learn

- Worker pool
- Pipeline
- Fan-out / fan-in
- Cancellation with done channel and context

## Concept

### Plain English

Production Go combines a few patterns:

**Worker pool:** fixed goroutines pull jobs from a channel, limiting parallelism.

**Pipeline:** stages connected by channels — each stage is a goroutine transforming a stream.

**Fan-out / fan-in:** split work across workers (fan-out), merge results (fan-in).

**Done channel:** `close(done)` broadcasts shutdown; workers select on `<-done` or use `context.Context`.

### Pattern diagram (ASCII)

```
Worker pool:

  jobs --> [worker]--\
  jobs --> [worker] ---> results
  jobs --> [worker]--/

Pipeline:

  in --> [square] --> [validate] --> out

Fan-in:

  ch1 --\
  ch2 ----> merge --> out
  ch3 --/
```

### Go syntax

```go
func worker(jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- process(j)
    }
}
```

## Annotated examples

```go
// WHY close(results) in separate goroutine after WaitGroup:
// workers must finish sending before close; avoid send on closed channel.
go func() {
    wg.Wait()
    close(results)
}()
```

## Common mistakes

- **Closing jobs before workers start:** Order matters — start workers, then send, then close jobs.
- **No cancellation in long loops:** Always select on ctx.Done().
- **Fan-in without WaitGroup:** Output closes too early, losing values.
- **Unbuffered pipeline deadlock:** Match stage speeds or use buffered channels.

## Further reading

- [Go blog — Pipelines and cancellation](https://go.dev/blog/pipelines)
- [Concurrency patterns (talk)](https://go.dev/talks/2012/concurrency.slide)

## API spec

| Function | Behavior |
|----------|----------|
| `ProcessJobs` | Worker pool; result = job * 2 |
| `Pipeline` | Square each value |
| `FanIn` | Merge until all inputs closed |
| `ProcessUntilDone` | Collect processed jobs until ctx/done/jobs closed |

## Before moving on

- [ ] I can sketch worker pool vs pipeline
- [ ] I close channels in the right order
- [ ] I handle cancellation with context or done
- [ ] All tests pass: `go test -v`

## Next module

When all concurrency lessons pass, continue to **05-tooling/01-testing**.
