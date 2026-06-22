# Go Learning Course

A self-contained, test-driven Go course. Each lesson gives you reading material, failing tests, and progressive hints. You learn by making tests pass.

## Prerequisites

- [Go 1.22+](https://go.dev/dl/) installed (`go version` should report 1.22 or later)
- Basic terminal comfort (`cd`, `ls`, running commands)
- A text editor or IDE

## Course map

| Module | Topics |
|--------|--------|
| **01-foundations** | Variables, functions, errors, slices/maps, structs |
| **02-interfaces** | Basics, composition, io contracts, type assertions, custom errors |
| **03-stdlib** | strings/fmt, file I/O, HTTP client, JSON, context |
| **04-concurrency** | Goroutines, channels, select, sync, patterns |
| **05-tooling** | Testing, modules, benchmarks, build tags |
| **06-projects** | CLI log parser, REST API, worker pool, capstone ideas |

## Quick start

Pick a lesson and run its tests:

```bash
cd 01-foundations/01-variables
go test -v
```

Tests fail on day one — that is expected. Open `README.md` in the same folder, read the concept, then edit `solution.go` until tests pass.

## Running tests

Each lesson is an independent Go module with its own `go.mod`. Always run tests from inside the lesson directory:

```bash
cd 01-foundations/01-variables
go test -v
```

Run all foundation lessons from the repo root (requires `go.work`):

```bash
go test ./01-foundations/01-variables ./01-foundations/02-functions \
  ./01-foundations/03-errors ./01-foundations/04-slices-maps \
  ./01-foundations/05-structs-methods
```

Or loop over every lesson in a module:

```bash
for d in 01-foundations/*/; do (cd "$d" && go test) || exit 1; done
for d in 02-interfaces/*/; do (cd "$d" && go test) || exit 1; done
```

Run all interface lessons from the repo root:

```bash
go test ./02-interfaces/01-basics ./02-interfaces/02-composition \
  ./02-interfaces/03-io-contracts ./02-interfaces/04-type-assertions \
  ./02-interfaces/05-custom-errors
```

Run all stdlib lessons from the repo root:

```bash
go test ./03-stdlib/01-strings-fmt ./03-stdlib/02-file-io \
  ./03-stdlib/03-http-client ./03-stdlib/04-json ./03-stdlib/05-context
```

Or loop:

```bash
for d in 03-stdlib/*/; do (cd "$d" && go test) || exit 1; done
for d in 04-concurrency/*/; do (cd "$d" && go test) || exit 1; done
```

Run all concurrency lessons from the repo root:

```bash
go test ./04-concurrency/01-goroutines ./04-concurrency/02-channels \
  ./04-concurrency/03-select ./04-concurrency/04-sync ./04-concurrency/05-patterns
```

Use race detection for this module:

```bash
go test -race ./04-concurrency/01-goroutines
```

Run all tooling lessons from the repo root:

```bash
go test ./05-tooling/01-testing ./05-tooling/02-modules \
  ./05-tooling/03-benchmarks ./05-tooling/04-build-flags
```

Or loop:

```bash
for d in 05-tooling/*/; do (cd "$d" && go test) || exit 1; done
```

Build-tags lesson also requires:

```bash
cd 05-tooling/04-build-flags
go test -tags=feature -run TestFeatureFlag_Enabled -v
```

Run all project lessons from the repo root:

```bash
go test ./06-projects/01-cli-tool ./06-projects/02-rest-api ./06-projects/03-worker-pool
```

## Race detection

When a lesson touches concurrency (Module 4 and some projects), use:

```bash
go test -race ./...
```

Module 1 does not require `-race`, but it is good to know the flag exists.

## Progression rules

1. Start at `01-foundations/01-variables` and work in order within each module.
2. **Do not advance** until every test in the current lesson passes (`go test` exits 0).
3. Read the lesson README fully before coding.
4. Do not edit `exercise_test.go` — implement everything in `solution.go`.

## Using hints without cheating yourself

Each lesson has a `hints.md` with three progressive hints separated by horizontal rules.

- Try for **at least 30 minutes** before opening hints.
- Read **only the first hint** initially. Implement, re-run tests, think again.
- Open the next hint only when you are truly stuck — not when you want a shortcut.
- The third hint is nearly a solution. Using it teaches you less. That is fine occasionally, but do not make it a habit.

## Module paths

Lessons use module paths like `github.com/yourname/go-course/foundations/variables`. If you fork this repo, search-and-replace `yourname` in each `go.mod` and test import path, or keep the placeholder — it does not affect local learning.

Physical directories use numeric prefixes (`01-variables`); Go **package names** drop the prefix (`package variables`).

## What each lesson contains

```
01-foundations/01-variables/
  README.md           # concept, examples, checklist
  exercise_test.go    # DO NOT EDIT — your spec
  hints.md            # unlockable help
  solution.go         # YOU implement here
  go.mod
```
