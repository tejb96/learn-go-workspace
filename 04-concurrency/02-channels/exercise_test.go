// DO NOT EDIT — implement the solution in solution.go

package channels_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/yourname/go-course/concurrency/channels"
)

func TestPingPong(t *testing.T) {
	if !channels.PingPong(1000) {
		t.Fatal("PingPong(1000) failed or timed out")
	}
}

func TestBufferedCollect(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	got := channels.BufferedCollect(ch, 3)
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("BufferedCollect() = %v, want %v", got, want)
	}
}

func TestSendWithTimeout_UnbufferedBlocks(t *testing.T) {
	ch := make(chan int) // unbuffered — no receiver yet
	ok := channels.SendWithTimeout(ch, 42, 20*time.Millisecond)
	if ok {
		t.Fatal("send should time out without receiver on unbuffered channel")
	}
}

func TestSendWithTimeout_BufferedSucceeds(t *testing.T) {
	ch := make(chan int, 1)
	if !channels.SendWithTimeout(ch, 1, 20*time.Millisecond) {
		t.Fatal("send should succeed on buffered channel with space")
	}
}
