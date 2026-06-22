// DO NOT EDIT — implement the solution in solution.go

package selectlesson_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/yourname/go-course/concurrency/selectlesson"
)

func TestFirstReady(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int, 1)
	ch2 <- 99

	got, err := selectlesson.FirstReady(ch1, ch2)
	if err != nil {
		t.Fatal(err)
	}
	if got != 99 {
		t.Fatalf("FirstReady = %d, want 99", got)
	}
}

func TestFirstReady_AllClosed(t *testing.T) {
	ch1 := make(chan int)
	close(ch1)
	ch2 := make(chan int)
	close(ch2)

	_, err := selectlesson.FirstReady(ch1, ch2)
	if !errors.Is(err, selectlesson.ErrNoReadyChannel) {
		t.Fatalf("err = %v, want ErrNoReadyChannel", err)
	}
}

func TestMerge(t *testing.T) {
	a := make(chan int, 2)
	b := make(chan int, 2)
	a <- 1
	a <- 2
	b <- 3
	close(a)
	close(b)

	out := selectlesson.Merge(a, b)
	var got []int
	for v := range out {
		got = append(got, v)
	}
	sortInts(got)
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Merge() = %v, want %v", got, want)
	}
}

func TestFirstReady_WaitsForFirst(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- 7
	}()

	got, err := selectlesson.FirstReady(ch1, ch2)
	if err != nil || got != 7 {
		t.Fatalf("got %d err %v, want 7 nil", got, err)
	}
}

func sortInts(s []int) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[i] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}
