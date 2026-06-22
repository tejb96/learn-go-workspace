// DO NOT EDIT — helpers for exercise_test.go (demonstrates t.Helper)

package testinglesson_test

import (
	"testing"

	"github.com/yourname/go-course/tooling/testinglesson"
)

func assertPalindrome(t *testing.T, s string, want bool) {
	t.Helper()
	got := testinglesson.IsPalindrome(s)
	if got != want {
		t.Fatalf("IsPalindrome(%q) = %v, want %v", s, got, want)
	}
}

func assertSumPositive(t *testing.T, nums []int, want int) {
	t.Helper()
	got := testinglesson.SumPositive(nums)
	if got != want {
		t.Fatalf("SumPositive(%v) = %d, want %d", nums, got, want)
	}
}
