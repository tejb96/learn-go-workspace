# select

## What you'll learn

- `select` for multiple channel operations
- Non-blocking ops with `default`
- Fan-in with select and goroutines

## Concept

### Plain English

**select** is like a switch for channels. It blocks until one of its cases can proceed, or runs `default` if none can (non-blocking).

Use it to wait on the first of several channels, time out, or cancel work.

### Go syntax

```go
select {
case v := <-ch1:
    use(v)
case ch2 <- x:
    sent()
case <-time.After(d):
    timeout()
default:
    noOp()
}
```

## Annotated examples

```go
// WHY select not sequential recv: sequential would starve ch2 if ch1 never sends.
select {
case v := <-ch1:
    return v, nil
case v := <-ch2:
    return v, nil
}
```

## Common mistakes

- **Empty select{}:** Blocks forever.
- **select in tight loop with default:** CPU spin — add sleep or rethink.
- **Not handling closed channels:** Receive on closed chan returns zero forever — use `v, ok := <-ch`.
- **Leaking goroutines in Merge:** Wait for all inputs before closing output.

## Further reading

- [Go spec — Select statements](https://go.dev/ref/spec#Select_statements)

## API spec

| Symbol | Behavior |
|--------|----------|
| `FirstReady(chs...)` | First available value; all closed empty → `ErrNoReadyChannel` |
| `Merge(chs...)` | Fan-in until all inputs closed |

## Before moving on

- [ ] I can write a blocking select on multiple channels
- [ ] I handle closed channels correctly
- [ ] All tests pass: `go test -v`
