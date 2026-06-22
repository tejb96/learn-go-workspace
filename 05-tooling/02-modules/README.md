# Go Modules and Workspaces

## What you'll learn

- `go mod init`, `go get`, `go mod tidy`
- `replace` for local modules
- `go work` for multi-module repos

## Concept

### Plain English

A **module** is a collection of Go packages versioned together. `go.mod` records the module path, Go version, and dependencies.

**`go mod init`** creates a new module. **`go get`** adds/upgrades dependencies. **`go mod tidy`** adds missing and removes unused requires.

**`replace`** swaps a module path with a local directory — common for developing two modules side by side.

**`go work`** (workspace) lets you work on multiple modules without publishing. `go.work` lists module directories with `use` directives. The repo root `go.work` does this for the whole course; this lesson has its own mini workspace.

### This lesson layout

```
02-modules/
  go.mod          # parent module with replace → ./examples/child
  go.work         # use . and ./examples/child
  examples/child/ # child module
  solution.go     # parse go.mod / go.work; call child
```

### Commands to try (read-only learning)

```bash
cd 05-tooling/02-modules
go work sync
go mod tidy
go test -v
```

```bash
# In a fresh dir elsewhere (do not run here):
go mod init example.com/foo
go get golang.org/x/text@latest
go mod tidy
```

### go.work diagram

```
go.work
  use ./                    → moduleslesson (this lesson)
  use ./examples/child      → moduleschild
```

## Annotated examples

```go
// go.mod replace — develop child locally without publishing.
replace github.com/yourname/go-course/tooling/moduleschild => ./examples/child
```

## Common mistakes

- **Editing go.sum by hand:** Let `go mod tidy` manage it.
- **Wrong replace path:** Relative to the go.mod file location.
- **Forgetting go.work when testing multi-module:** `go work sync` refreshes deps.
- **Module path mismatch:** Import path must match `module` directive.

## Further reading

- [Go modules reference](https://go.dev/ref/mod)
- [Go workspaces](https://go.dev/doc/tutorial/workspaces)

## API spec

| Function | Behavior |
|----------|----------|
| `ModulePath(contents)` | Parse first `module` line |
| `HasRequire(contents, path)` | true if path in require block |
| `WorkUsePaths(contents)` | Parse `use` paths in order |
| `ChildGreet(name)` | Delegate to `child.Greet` |

## Before moving on

- [ ] I read go.mod, go.work, and examples/child/go.mod
- [ ] I ran `go work sync` and `go test`
- [ ] I understand replace vs require
- [ ] All tests pass: `go test -v`
