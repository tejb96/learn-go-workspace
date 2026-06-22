# Capstone Projects

You have completed the structured lessons. These open-ended projects apply what you learned to **cloud and DevOps** tooling. There are no tests — you define success criteria.

Try for at least a few hours on one idea before jumping to the next.

---

## 1. Kubernetes pod health checker

Build a CLI that:

- Accepts a kubeconfig path and namespace
- Lists pods and reports which are not `Running` / not ready
- Outputs JSON or a table for CI pipelines

**Stretch goals:** watch mode, exit code non-zero when unhealthy pods exist, timeout per API call with `context`.

**Skills used:** CLI flags, JSON output, HTTP client, context, error wrapping.

**Resources:** [client-go](https://pkg.go.dev/k8s.io/client-go) (optional) or raw API with `net/http` + bearer token from kubeconfig.

---

## 2. Terraform plan summary parser

Build a tool that:

- Reads `terraform plan -out=plan.bin` JSON output (or `terraform show -json plan.bin`) from stdin or file
- Summarizes resources to **create**, **update**, **delete**
- Prints a one-screen summary for PR comments

**Stretch goals:** filter by resource type, fail if delete count exceeds threshold (safety gate in CI).

**Skills used:** JSON decoding, structs with tags, stdin/file input, table formatting.

**Resources:** [Terraform JSON output](https://developer.hashicorp.com/terraform/internals/json-format)

---

## 3. Metrics sidecar / log forwarder

Build a small HTTP service that:

- Exposes `GET /metrics` in Prometheus text format (counters you increment)
- Accepts `POST /events` with JSON log lines and buffers them
- Flushes buffer to stdout or a file every N seconds or on `SIGTERM`

**Stretch goals:** graceful shutdown with `context`, worker pool for flush, `go test` with `httptest`.

**Skills used:** HTTP server, concurrency, channels, signal handling, sync.

**Resources:** [Prometheus exposition format](https://prometheus.io/docs/instrumenting/exposition_formats/)

---

## Before you call it done

For whichever project you choose:

- [ ] `go test ./...` covers core logic (even if you design the tests yourself)
- [ ] README with build/run instructions
- [ ] Error handling on I/O and network paths
- [ ] No hard-coded secrets — env vars or flags

## Where to go next

- [Go official docs](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- Contribute back to this course or share your capstone repo

Congratulations on finishing the course modules.
