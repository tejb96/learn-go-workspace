// DO NOT EDIT — implement the solution in solution.go

package goroutines_test

import (
	"sync"
	"testing"

	"github.com/yourname/go-course/concurrency/goroutines"
)

func runConcurrentIncrements(inc func(), goroutines, each int, wg *sync.WaitGroup) {
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < each; j++ {
				inc()
			}
		}()
	}
}

func TestSafeCounter_Concurrent(t *testing.T) {
	var c goroutines.SafeCounter
	const goroutines = 50
	const each = 100
	var wg sync.WaitGroup
	runConcurrentIncrements(c.Inc, goroutines, each, &wg)
	wg.Wait()

	want := goroutines * each
	if got := c.Value(); got != want {
		t.Fatalf("SafeCounter.Value() = %d, want %d", got, want)
	}
}

func TestUnsafeCounter_OftenLosesUpdates(t *testing.T) {
	var c goroutines.UnsafeCounter
	const goroutines = 50
	const each = 100
	var wg sync.WaitGroup
	runConcurrentIncrements(c.Inc, goroutines, each, &wg)
	wg.Wait()

	want := goroutines * each
	got := c.Value()
	if got == want {
		t.Logf("UnsafeCounter reached %d (lucky or single-threaded); run go test -race to see data races", got)
	}
	// Educational: unsafe path is not required to pass a specific value.
}

func TestRunAll(t *testing.T) {
	var mu sync.Mutex
	n := 0
	fns := make([]func(), 10)
	for i := range fns {
		fns[i] = func() {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}
	goroutines.RunAll(fns)
	if n != 10 {
		t.Fatalf("n = %d, want 10", n)
	}
}
