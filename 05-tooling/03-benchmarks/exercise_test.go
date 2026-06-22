// DO NOT EDIT — implement the solution in solution.go

package benchmarks_test

import (
	"testing"

	"github.com/yourname/go-course/tooling/benchmarks"
)

func TestFib(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 0, want: 0},
		{n: 1, want: 1},
		{n: 10, want: 55},
		{n: 20, want: 6765},
	}

	for _, tt := range tests {
		if got := benchmarks.Fib(tt.n); got != tt.want {
			t.Fatalf("Fib(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarks.Fib(20)
	}
}

func BenchmarkFibCached(b *testing.B) {
	cache := make(map[int]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarks.FibCached(20, cache)
	}
}
