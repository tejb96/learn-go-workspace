# Channels

## What you'll learn

- Unbuffered vs buffered channels
- Blocking semantics
- Channel direction in types (`<-chan`, `chan<-`)

## Concept

### Plain English

**Channels** pass values between goroutines. An **unbuffered** channel has no queue — a send blocks until a receive happens (and vice versa). Handshake synchronization.

A **buffered** channel has capacity `n`. Sends proceed without a receiver until the buffer is full; receives proceed until empty.

Learn unbuffered first — they teach you why goroutines block and how to reason about coordination.

### Blocking semantics (ASCII)

**Unbuffered — rendezvous:**

```
  sender goroutine          channel (cap 0)          receiver goroutine
        |                        |                          |
        |---- send v ------------>|                          |
        |      (blocks)           |                          |
        |                         |<------- recv ------------|
        |      (unblocks)         |        (unblocks)        |
```

**Buffered — queue:**

```
  make(chan int, 3)

  [ 1 | 2 | _ ]   send 3 → ok (no receiver yet)
  [ 1 | 2 | 3 ]   send 4 → blocks until recv frees slot
```

### Go syntax

```go
ch := make(chan int)     // unbuffered
buf := make(chan int, 5) // buffered, cap 5

ch <- 42    // send
v := <-ch   // receive
close(ch)
```

## Annotated examples

```go
// WHY unbuffered for PingPong: forces strict alternation — each send pairs with recv.
ping := make(chan struct{})
```

```go
// WHY buffered for producer burst: decouple short bursts without blocking producer immediately.
jobs := make(chan Job, 100)
```

## Common mistakes

- **Sending on nil channel:** Blocks forever.
- **Closing from receiver side:** Only sender should close (usually).
- **Assuming buffered = async:** Still blocks when full.
- **Range over channel without close:** Deadlock if sender never closes.

## Further reading

- [Go spec — Channel types](https://go.dev/ref/spec#Channel_types)
- [Effective Go — Channels](https://go.dev/doc/effective_go#channels)

## API spec

| Function | Behavior |
|----------|----------|
| `PingPong(n)` | Unbuffered token pass n times; return true |
| `BufferedCollect(ch, count)` | Receive `count` ints in order |
| `SendWithTimeout(ch, v, d)` | false if send blocks longer than d |

## Before moving on

- [ ] I can draw unbuffered vs buffered blocking
- [ ] I know cap 0 means unbuffered
- [ ] All tests pass: `go test -v`
