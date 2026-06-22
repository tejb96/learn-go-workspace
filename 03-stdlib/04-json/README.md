# JSON Encoding

## What you'll learn

- `json.Marshal` / `json.Unmarshal`
- Struct tags: names, `omitempty`, ignore with `-`
- Unknown JSON fields and API-safe views

## Concept

### Plain English

Go structs map to JSON objects via **struct tags**. The encoder uses exported field names by default, but tags customize output: rename fields, omit zero values, or exclude secrets entirely.

When decoding, extra JSON fields are **ignored** by default — forward-compatible API evolution.

### Go syntax

```go
type Profile struct {
    Name     string `json:"name"`
    Age      int    `json:"age,omitempty"`
    Password string `json:"-"`
}
```

```go
data, err := json.MarshalIndent(v, "", "  ")
err = json.Unmarshal(data, &dest)
```

## Annotated examples

```go
// WHY omitempty: keeps payloads small; zero age means "not set" in API.
Age int `json:"age,omitempty"`
```

```go
// WHY json:"-": password never appears in marshaled output, even accidentally.
Password string `json:"-"`
```

## Common mistakes

- **Unexported fields:** JSON ignores lowercase field names.
- **Unmarshal into value not pointer:** Must pass `&dest`.
- **Expecting null vs missing:** `omitempty` omits zero values on marshal; unmarshaling missing keys leaves zero values.
- **Using map for everything:** Structs give type safety; maps for truly dynamic JSON.

## Further reading

- [Package encoding/json](https://pkg.go.dev/encoding/json)
- [JSON and Go](https://go.dev/blog/json)

## API spec

| Symbol | Behavior |
|--------|----------|
| `ToJSON(v)` | Indented JSON, 2 spaces |
| `FromJSON(data, dest)` | Unmarshal into dest |
| `PublicView(p)` | Name and Age only |

## Before moving on

- [ ] I use struct tags for names and omitempty
- [ ] I exclude sensitive fields with `-`
- [ ] All tests pass: `go test -v`
