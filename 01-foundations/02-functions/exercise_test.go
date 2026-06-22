// DO NOT EDIT — implement the solution in solution.go

package functions_test

import (
	"errors"
	"testing"

	"github.com/yourname/go-course/foundations/functions"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "empty", nums: nil, want: 0},
		{name: "single", nums: []int{5}, want: 5},
		{name: "multiple", nums: []int{1, 2, 3, 4}, want: 10},
		{name: "negatives", nums: []int{-1, 1}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := functions.Sum(tt.nums...)
			if got != tt.want {
				t.Fatalf("Sum(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr error
	}{
		{name: "even division", a: 10, b: 2, want: 5},
		{name: "truncates toward zero", a: 7, b: 2, want: 3},
		{name: "negative", a: -7, b: 2, want: -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := functions.Divide(tt.a, tt.b)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("Divide(%d, %d) error = %v, want %v", tt.a, tt.b, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("Divide(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide_DivideByZero(t *testing.T) {
	_, err := functions.Divide(1, 0)
	if err == nil {
		t.Fatal("Divide(1, 0) expected error, got nil")
	}
}
