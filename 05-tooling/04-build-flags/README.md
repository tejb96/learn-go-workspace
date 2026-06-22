# Build Tags and Flags

## What you'll learn

- `//go:build` constraints (Go 1.17+)
- Building with `-tags`
- Intro to `-ldflags` (link-time variables)

## Concept

### Plain English

**Build tags** select which files compile. Put a constraint at the top of a file:

```go
//go:build feature

package buildflags
```

Only builds when you pass `-tags=feature`. The negated form `//go:build !feature` means "build when feature tag is **absent**."

This replaces the old `// +build` comment syntax.

**`-ldflags`** sets string variables at link time (common for version injection). This lesson focuses on tags; ldflags are mentioned for awareness.

### Go syntax

```bash
go test -v                      # default: feature_off.go
go test -tags=feature -v        # feature_on.go
go build -tags=feature .
```

### File layout in this lesson

| File | Build tag | FeatureFlag |
|------|-----------|-------------|
| `feature_off.go` | `!feature` | `"off"` |
| `feature_on.go` | `feature` | you set `"on"` |
| `solution.go` | always | `DebugMode()` |

## Annotated examples

```go
// WHY separate files: clean compile-time feature split without runtime if chains.
//go:build feature
```

## Common mistakes

- **Tag on wrong line:** Must be immediately above `package`, only `//go:build` (not block comment).
- **Forgetting -tags when testing:** `TestFeatureFlag_Enabled` requires `-tags=feature`.
- **Overlapping tags:** Two files both match → duplicate symbol error.
- **Using tags for runtime config:** Use env/config instead; tags need recompile.

## Further reading

- [Build constraints](https://pkg.go.dev/cmd/go#hdr-Build_constraints)
- [Go 1.17 build tags](https://go.dev/doc/go1.17#go-build)

## API spec

| Symbol | Behavior |
|--------|----------|
| `FeatureFlag()` | `"off"` default build; `"on"` with `-tags=feature` |
| `DebugMode()` | `"default"` |

## Before moving on

- [ ] `go test -v` passes (default build)
- [ ] `go test -tags=feature -run TestFeatureFlag_Enabled -v` passes
- [ ] I understand `//go:build` vs `-tags`

## Next module

When all tooling lessons pass, continue to **06-projects/01-cli-tool**.
