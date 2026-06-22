// DO NOT EDIT — implement the solution in solution.go

package composition_test

import (
	"reflect"
	"testing"

	"github.com/yourname/go-course/interfaces/composition"
)

// mockReader implements Reader only.
type mockReader struct {
	data []byte
}

func (m *mockReader) Read() []byte { return m.data }

// mockWriter implements Writer only.
type mockWriter struct {
	stored []byte
}

func (m *mockWriter) Write(p []byte) int {
	m.stored = append(m.stored, p...)
	return len(p)
}

// mockReadWriter implements ReadWriter via embedding-style dual methods.
type mockReadWriter struct {
	mockReader
	mockWriter
}

func TestDescribe(t *testing.T) {
	tests := []struct {
		name string
		v    any
		want string
	}{
		{name: "reader only", v: &mockReader{data: []byte("x")}, want: "reader"},
		{name: "writer only", v: &mockWriter{}, want: "writer"},
		{name: "readwriter", v: &mockReadWriter{mockReader: mockReader{data: []byte("x")}}, want: "readwriter"},
		{name: "unrelated", v: 42, want: "none"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := composition.Describe(tt.v)
			if got != tt.want {
				t.Fatalf("Describe() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDuplicate(t *testing.T) {
	r := &mockReader{data: []byte("go")}
	w1 := &mockWriter{}
	w2 := &mockWriter{}

	got := composition.Duplicate(r, w1, w2)
	want := []int{2, 2}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Duplicate() = %v, want %v", got, want)
	}
	if string(w1.stored) != "go" || string(w2.stored) != "go" {
		t.Fatalf("writers did not receive data: w1=%q w2=%q", w1.stored, w2.stored)
	}
}

func TestDuplicate_NoWriters(t *testing.T) {
	r := &mockReader{data: []byte("x")}
	got := composition.Duplicate(r)
	if got == nil || len(got) != 0 {
		t.Fatalf("Duplicate() with no writers = %v, want empty slice", got)
	}
}
