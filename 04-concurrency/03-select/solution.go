package selectlesson

import "errors"

var ErrNoReadyChannel = errors.New("no ready channel")

// FirstReady returns the value from the first channel that can send without blocking.
// If all channels are closed and empty, return ErrNoReadyChannel.
func FirstReady(chs ...<-chan int) (int, error) {
	return 0, nil
}

// Merge returns a channel receiving from all inputs until all are drained and closed.
func Merge(chs ...<-chan int) <-chan int {
	ch := make(chan int)
	close(ch)
	return ch
}
