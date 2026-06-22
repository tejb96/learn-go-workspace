// DO NOT EDIT — implement the solution in solution.go

package patterns_test

import (
	"context"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/yourname/go-course/concurrency/patterns"
)

func TestProcessJobs(t *testing.T) {
	jobs := make(chan int, 5)
	for _, j := range []int{1, 2, 3, 4, 5} {
		jobs <- j
	}
	close(jobs)

	out := patterns.ProcessJobs(jobs, 2)
	var got []int
	for r := range out {
		got = append(got, r)
	}
	sort.Ints(got)
	want := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ProcessJobs() = %v, want %v", got, want)
	}
}

func TestPipeline(t *testing.T) {
	in := make(chan int, 3)
	in <- 2
	in <- 3
	in <- 4
	close(in)

	out := patterns.Pipeline(in)
	var got []int
	for v := range out {
		got = append(got, v)
	}
	sort.Ints(got)
	want := []int{4, 9, 16}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Pipeline() = %v, want %v", got, want)
	}
}

func TestFanIn(t *testing.T) {
	a := make(chan int, 1)
	b := make(chan int, 1)
	a <- 1
	b <- 2
	close(a)
	close(b)

	out := patterns.FanIn(a, b)
	var got []int
	for v := range out {
		got = append(got, v)
	}
	sort.Ints(got)
	if !reflect.DeepEqual(got, []int{1, 2}) {
		t.Fatalf("FanIn() = %v", got)
	}
}

func TestProcessUntilDone_CancelViaDone(t *testing.T) {
	jobs := make(chan int, 100)
	for i := 1; i <= 100; i++ {
		jobs <- i
	}
	done := make(chan struct{})
	close(done)

	got := patterns.ProcessUntilDone(context.Background(), jobs, done, func(n int) int {
		return n * 10
	})
	if len(got) != 0 {
		t.Fatalf("expected no results when done closed immediately, got %v", got)
	}
}

func TestProcessUntilDone_ContextCancel(t *testing.T) {
	jobs := make(chan int)
	go func() {
		for i := 1; ; i++ {
			jobs <- i
			time.Sleep(time.Millisecond)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	done := make(chan struct{})

	got := patterns.ProcessUntilDone(ctx, jobs, done, func(n int) int { return n })
	if len(got) == 0 {
		t.Fatal("expected some results before cancel")
	}
}

func TestProcessUntilDone_ProcessesJobs(t *testing.T) {
	jobs := make(chan int, 3)
	jobs <- 1
	jobs <- 2
	jobs <- 3
	close(jobs)
	done := make(chan struct{})

	got := patterns.ProcessUntilDone(context.Background(), jobs, done, func(n int) int {
		return n + 1
	})
	sort.Ints(got)
	want := []int{2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}
