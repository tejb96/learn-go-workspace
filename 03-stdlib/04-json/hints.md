Try for at least 30 minutes before reading this.

## Hint 1

Struct tags control JSON field names: `json:"name"`, `json:"age,omitempty"`, `json:"-"` to skip. `encoding/json` ignores unknown fields by default when unmarshaling into structs.

---

## Hint 2

`ToJSON` uses `json.MarshalIndent(v, "", "  ")`. `FromJSON` is `json.Unmarshal`. `PublicView` copies only exported-safe fields — clear Password and Internal manually.

---

## Hint 3

```go
func ToJSON(v any) ([]byte, error) {
    return json.MarshalIndent(v, "", "  ")
}

func FromJSON(data []byte, dest any) error {
    return json.Unmarshal(data, dest)
}

func PublicView(p Profile) Profile {
    return Profile{Name: p.Name, Age: p.Age}
}
```

The `-` tag excludes Password from marshaling automatically; `PublicView` is for explicit API shaping.
