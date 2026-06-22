// DO NOT EDIT — implement the solution in solution.go

package testinglesson_test

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "empty", s: "", want: true},
		{name: "single", s: "a", want: true},
		{name: "classic", s: "racecar", want: true},
		{name: "not palindrome", s: "go", want: false},
		{name: "case sensitive", s: "Abba", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertPalindrome(t, tt.s, tt.want)
		})
	}
}

func TestSumPositive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "mixed", nums: []int{-1, 2, 0, 3}, want: 5},
		{name: "all non-positive", nums: []int{-1, 0}, want: 0},
		{name: "empty", nums: nil, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertSumPositive(t, tt.nums, tt.want)
		})
	}
}
