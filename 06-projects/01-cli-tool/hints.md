Try for at least 30 minutes before reading this.

## Hint 1

Use the `flag` package bound to a custom FlagSet to parse `-level`, `-format`, and `-file`. Read lines with `bufio.Scanner`. Plain output: one line per entry; JSON: `json.Marshal` the slice.

---

## Hint 2

`ParseLine`: `strings.Fields` or split on first space — first token is level, remainder is message. `FilterByLevel`: compare `strings.EqualFold`. If `-file` set, `os.Open`; else use `stdin`.

---

## Hint 3

```go
fs := flag.NewFlagSet("logparser", flag.ContinueOnError)
level := fs.String("level", "", "filter level")
format := fs.String("format", "plain", "plain or json")
file := fs.String("file", "", "log file path")
fs.Parse(args)

var r io.Reader = stdin
if *file != "" {
    f, err := os.Open(*file)
    // ...
    r = f
}
```

Output JSON array even for a single entry.
