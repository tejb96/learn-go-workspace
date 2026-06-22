# CLI Log Parser

## What you'll learn

- Building a small CLI with the `flag` package
- Reading from stdin or a file
- Formatting output as plain text or JSON

## Project brief

Build a log parser that:

1. Reads log lines from **stdin** or **`-file`**
2. Parses lines as `LEVEL message...`
3. Filters by **`-level`** (case-insensitive; omit to include all)
4. Writes **`-format=plain`** or **`-format=json`** to stdout

Example input:

```
INFO server started
ERROR disk full
WARN low memory
```

```bash
go test -v
# when passing:
echo -e "INFO hi\nERROR bye" | go run . -level=ERROR -format=plain
```

Implement in `solution.go`. Optional `main.go` for manual runs:

```go
func main() {
    if err := clitool.Run(os.Args[1:], os.Stdin, os.Stdout); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
```

## Common mistakes

- **Not trimming empty lines:** Skip blank lines when scanning.
- **JSON output as single object:** Tests expect a JSON **array** of entries.
- **Flag parsing on global flag.CommandLine:** Use `flag.NewFlagSet` for testable `Run`.

## Further reading

- [Package flag](https://pkg.go.dev/flag)
- [Package bufio](https://pkg.go.dev/bufio)

## Before moving on

- [ ] `go test -v` passes
- [ ] I tried the CLI manually with stdin and `-file`
- [ ] Plain and JSON formats both work
