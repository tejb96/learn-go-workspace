// DO NOT EDIT — implement the solution in solution.go

package slicesmaps_test

import (
	"reflect"
	"testing"

	"github.com/yourname/go-course/foundations/slicesmaps"
)

func TestAppendUnique(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		v    int
		want []int
	}{
		{name: "append new", s: []int{1, 2}, v: 3, want: []int{1, 2, 3}},
		{name: "skip duplicate", s: []int{1, 2}, v: 2, want: []int{1, 2}},
		{name: "empty slice", s: nil, v: 1, want: []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slicesmaps.AppendUnique(tt.s, tt.v)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("AppendUnique(%v, %d) = %v, want %v", tt.s, tt.v, got, tt.want)
			}
		})
	}
}

func TestAppendUnique_DoesNotMutateSharedBackingArrayUnexpectedly(t *testing.T) {
	orig := []int{1, 2}
	shared := orig[:2:2] // len=2 cap=2 — append must reallocate
	got := slicesmaps.AppendUnique(shared, 3)
	_ = got
	if len(orig) != 2 || orig[0] != 1 || orig[1] != 2 {
		t.Fatalf("orig mutated = %v, want [1 2]", orig)
	}
}

func TestSafeSlice(t *testing.T) {
	s := []int{10, 20, 30, 40}

	tests := []struct {
		name     string
		low, high int
		want     []int
	}{
		{name: "middle", low: 1, high: 3, want: []int{20, 30}},
		{name: "clamp high", low: 0, high: 100, want: []int{10, 20, 30, 40}},
		{name: "clamp low", low: -5, high: 2, want: []int{10, 20}},
		{name: "empty when low equals high", low: 2, high: 2, want: []int{}},
		{name: "empty when low greater than high", low: 3, high: 1, want: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slicesmaps.SafeSlice(s, tt.low, tt.high)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("SafeSlice(...) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapInvert(t *testing.T) {
	tests := []struct {
		name string
		in   map[string]int
		want map[int]string
	}{
		{
			name: "basic",
			in:   map[string]int{"a": 1, "b": 2},
			want: map[int]string{1: "a", 2: "b"},
		},
		{
			name: "duplicate values last wins",
			in:   map[string]int{"first": 1, "second": 1},
			want: map[int]string{1: "second"},
		},
		{
			name: "empty map",
			in:   map[string]int{},
			want: map[int]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slicesmaps.MapInvert(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("MapInvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapInvert_NilMapRead(t *testing.T) {
	var m map[string]int
	// Reading nil map is safe; ranging is zero iterations.
	got := slicesmaps.MapInvert(m)
	if got == nil || len(got) != 0 {
		t.Fatalf("MapInvert(nil) = %v, want empty map", got)
	}
}

func TestCapAfterAppend(t *testing.T) {
	s := make([]int, 0, 2) // len 0 cap 2
	got := slicesmaps.CapAfterAppend(s)
	// After one append on cap=2, cap may stay 2 or grow depending on len;
	// with len=0 cap=2, first append fills to len=1 cap=2 still.
	if got != 2 {
		t.Fatalf("CapAfterAppend() = %d, want 2", got)
	}

	s2 := []int{1, 2} // len 2, cap 2
	got2 := slicesmaps.CapAfterAppend(s2)
	// append triggers growth, typically double: 4
	if got2 < 4 {
		t.Fatalf("CapAfterAppend(full slice) = %d, want at least 4", got2)
	}
}
