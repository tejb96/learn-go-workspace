package synclesson

// ParallelMap applies fn to each element of in using at most workers goroutines.
// Results preserve input order.
func ParallelMap(in []int, fn func(int) int, workers int) []int {
	return nil
}

// OnceValue returns a function that calls fn exactly once and caches the result.
func OnceValue(fn func() int) func() int {
	return func() int { return 0 }
}
