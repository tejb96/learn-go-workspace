Try for at least 30 minutes before reading this.

## Hint 1

This lesson includes a **local go.work** and `examples/child/` module. Read the README for `go mod init`, `go get`, `go mod tidy`, and `go work`. Code parses `go.mod` / `go.work` text — use `strings` or regex, not the go command.

---

## Hint 2

`ModulePath`: find line starting with `module `, return trimmed path. `HasRequire`: search for `modulePath` in a line containing `require`. `WorkUsePaths`: lines inside `use (` block starting with `./` or `.`.

---

## Hint 3

```go
import "github.com/yourname/go-course/tooling/moduleschild"

func ChildGreet(name string) string {
    return child.Greet(name)
}
```

For `WorkUsePaths`, split lines, track `inUse` after `use (`, collect trimmed paths until `)`.

Run from this directory: `go work sync && go test -v`
