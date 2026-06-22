package benchmarks

// Fib returns the nth Fibonacci number (0-indexed: Fib(0)=0, Fib(1)=1).
func Fib(n int) int {
	return -1
}

// FibCached returns Fib using a map cache for repeated calls in benchmarks.
func FibCached(n int, cache map[int]int) int {
	return -1
}
