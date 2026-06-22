package patterns

import "context"

// ProcessJobs reads jobs from jobs, processes each (multiply by 2) with workers goroutines,
// and sends results on the returned channel. Close results when all jobs processed.
func ProcessJobs(jobs <-chan int, workers int) <-chan int {
	ch := make(chan int)
	close(ch)
	return ch
}

// Pipeline squares each input value.
func Pipeline(in <-chan int) <-chan int {
	ch := make(chan int)
	close(ch)
	return ch
}

// FanIn merges multiple input channels into one output channel until all inputs close.
func FanIn(chs ...<-chan int) <-chan int {
	ch := make(chan int)
	close(ch)
	return ch
}

// ProcessUntilDone runs processor on jobs until done is closed or ctx is canceled.
func ProcessUntilDone(ctx context.Context, jobs <-chan int, done <-chan struct{}, processor func(int) int) []int {
	return nil
}
