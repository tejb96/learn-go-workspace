Try for at least 30 minutes before reading this.

## Hint 1

Build tags are line comments before the package clause: `//go:build feature`. Files with conflicting tags compile exclusively. Default build uses `feature_off.go`; `-tags=feature` uses `feature_on.go`.

---

## Hint 2

Edit `feature_on.go` to return `"on"`. `feature_off.go` already returns `"off"`. `DebugMode` in `solution.go` returns `"default"`.

---

## Hint 3

Verify both builds:

```bash
go test -v
go test -tags=feature -v -run TestFeatureFlag
```

Optional ldflags (README only): `go build -ldflags "-X main.version=1.0"` — not required for tests.
