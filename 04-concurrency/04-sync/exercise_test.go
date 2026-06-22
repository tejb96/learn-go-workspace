// DO NOT EDIT — implement the solution in solution.go

package synclesson_test

import (
	"reflect"
	"sync/atomic"
	"testing"

	"github.com/yourname/go-course/concurrency/synclesson"
)

func TestParallelMap(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	got := synclesson.ParallelMap(in, func(n int) int { return n * n }, 2)
	want := []int{1, 4, 9, 16, 25}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParallelMap() = %v, want %v", got, want)
	}
}

func TestParallelMap_Empty(t *testing.T) {
	got := synclesson.ParallelMap(nil, func(n int) int { return n }, 4)
	if got == nil || len(got) != 0 {
		t.Fatalf("ParallelMap(nil) = %v, want empty slice", got)
	}
}

func TestParallelMap_WorkerLimit(t *testing.T) {
	var concurrent int32
	var peak int32
	in := make([]int, 20)
	for i := range in {
		in[i] = i
	}
	_ = synclesson.ParallelMap(in, func(n int) int {
		cur := atomic.AddInt32(&concurrent, 1)
		for {
			old := atomic.LoadInt32(&peak)
			if cur <= old || atomic.CompareAndSwapInt32(&peak, old, cur) {
				break
			}
		}
		atomic.AddInt32(&concurrent, -1)
		return n
	}, 3)
	if peak > 3 {
		t.Fatalf("peak concurrent = %d, want at most 3", peak)
	}
}

func TestOnceValue(t *testing.T) {
	var calls int32
	fn := synclesson.OnceValue(func() int {
		atomic.AddInt32(&calls, 1)
		return 42
	})
	if fn() != 42 || fn() != 42 {
		t.Fatal("OnceValue should return cached 42")
	}
	if atomic.LoadInt32(&calls) != 1 {
		t.Fatalf("fn called %d times, want 1", calls)
	}
}
