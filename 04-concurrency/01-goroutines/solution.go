package goroutines

// UnsafeCounter increments without synchronization — data races when used concurrently.
// Run: go test -race ./... to see the race detector report issues.
type UnsafeCounter struct {
	n int
}

func (c *UnsafeCounter) Inc() {
	c.n++
}

func (c *UnsafeCounter) Value() int {
	return c.n
}

// SafeCounter increments with a mutex for concurrent use.
type SafeCounter struct {
	// TODO: add sync.Mutex field
	n int
}

func (c *SafeCounter) Inc() {
}

func (c *SafeCounter) Value() int {
	return 0
}

// RunAll starts fn in a goroutine for each item and waits for all to finish.
func RunAll(fns []func()) {
}
