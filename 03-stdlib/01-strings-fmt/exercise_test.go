// DO NOT EDIT — implement the solution in solution.go

package stringsfmt_test

import (
	"testing"

	"github.com/yourname/go-course/stdlib/stringsfmt"
)

func TestSlugify(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "simple", in: "Hello World", want: "hello-world"},
		{name: "trim and underscores", in: "  Go_is_fun  ", want: "go-is-fun"},
		{name: "collapse hyphens", in: "foo--bar", want: "foo-bar"},
		{name: "empty", in: "", want: ""},
		{name: "unicode preserved", in: "Café Latte", want: "café-latte"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringsfmt.Slugify(tt.in)
			if got != tt.want {
				t.Fatalf("Slugify(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestFormatTable(t *testing.T) {
	got := stringsfmt.FormatTable(
		[]string{"name", "age"},
		[][]string{{"Ada", "36"}, {"Lin", "28"}},
	)
	want := "name | age\nAda | 36\nLin | 28\n"
	if got != want {
		t.Fatalf("FormatTable() = %q, want %q", got, want)
	}
}

func TestFormatTable_EmptyRows(t *testing.T) {
	got := stringsfmt.FormatTable([]string{"x"}, nil)
	want := "x\n"
	if got != want {
		t.Fatalf("FormatTable() = %q, want %q", got, want)
	}
}

func TestJoinLines(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want string
	}{
		{name: "multiple", in: []string{"a", "b"}, want: "a\nb"},
		{name: "single", in: []string{"only"}, want: "only"},
		{name: "empty slice", in: []string{}, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringsfmt.JoinLines(tt.in)
			if got != tt.want {
				t.Fatalf("JoinLines() = %q, want %q", got, tt.want)
			}
		})
	}
}
